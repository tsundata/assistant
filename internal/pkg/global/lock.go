package global

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"go.etcd.io/etcd/client/v3/concurrency"
	"strings"
	"sync"
	"time"
)

type Locker struct {
	client *etcd.Client
	op     *Operation
}

func NewLocker(client *etcd.Client) *Locker {
	return &Locker{client: client, op: &Operation{
		AcquireTimeout: 5 * time.Second,
		TTL:            10 * time.Second,
		Try:            true,
	}}
}

type Operation struct {
	AcquireTimeout time.Duration
	TTL            time.Duration
	Try            bool
}

type LockerOpt func(op *Operation)

func WithTry(try bool) LockerOpt {
	return LockerOpt(func(op *Operation) {
		op.Try = try
	})
}

func WithAcquireTimeout(timeout time.Duration) LockerOpt {
	return LockerOpt(func(op *Operation) {
		op.AcquireTimeout = timeout
	})
}

func WithLockTTL(ttl time.Duration) LockerOpt {
	return LockerOpt(func(op *Operation) {
		op.TTL = ttl
	})
}

func (l *Locker) Acquire(key string, opts ...LockerOpt) (*Lock, error) {
	op := &Operation{
		AcquireTimeout: l.op.AcquireTimeout,
		TTL:            l.op.TTL,
		Try:            l.op.Try,
	}

	for _, opt := range opts {
		opt(op)
	}

	ttl := int(op.TTL / time.Second)
	session, err := concurrency.NewSession(l.client.Client, concurrency.WithTTL(ttl))
	if err != nil {
		return nil, err
	}

	key = addPrefix(key)
	mutex := concurrency.NewMutex(session, key)

	if op.Try {
		tryLockErr := l.tryLock(mutex, op.AcquireTimeout)
		if errors.Is(tryLockErr, concurrency.ErrLocked) {
			err = session.Close()
			if err != nil {
				return nil, err
			}
			return nil, errors.New("already locked")
		}

		if tryLockErr != nil {
			err = session.Close()
			if err != nil {
				return nil, err
			}
			return nil, tryLockErr
		}
	} else {
		tryLockErr := l.lock(mutex, op.AcquireTimeout)

		if errors.Is(tryLockErr, context.DeadlineExceeded) {
			err = session.Close()
			if err != nil {
				return nil, err
			}
			return nil, errors.New("lock deadline exceeded")
		}

		if tryLockErr != nil {
			err = session.Close()
			if err != nil {
				return nil, err
			}
			return nil, tryLockErr
		}
	}

	session.Orphan()
	lock := &Lock{mutex: mutex, Mutex: &sync.Mutex{}, session: session}

	return lock, nil
}

func (l *Locker) lock(mutex *concurrency.Mutex, tryLockTimeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), tryLockTimeout)
	defer cancel()
	return mutex.Lock(ctx)
}

func (l *Locker) tryLock(mutex *concurrency.Mutex, tryLockTimeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), tryLockTimeout)
	defer cancel()
	return mutex.TryLock(ctx)
}

type Lock struct {
	*sync.Mutex
	mutex   *concurrency.Mutex
	session *concurrency.Session
}

func (l *Lock) Release() error {
	if l == nil {
		return errors.New("not found locked")
	}
	l.Lock()
	defer l.Unlock()

	defer func() {
		_ = l.session.Close()
	}()
	return l.mutex.Unlock(context.Background())
}

const prefix = "/etcd-lock"

func addPrefix(key string) string {
	if !strings.HasPrefix(key, "/") {
		key = "/" + key
	}
	return prefix + key
}
