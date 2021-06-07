package influx

import (
	"context"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"path"
	"time"
)

func UnaryServerInterceptor(in influxdb.Client, org string, bucket string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		startTime := time.Now()
		resp, err = handler(ctx, req)
		writePoint(startTime, in, org, bucket, info.FullMethod, err)
		return resp, err
	}
}

func StreamServerInterceptor(in influxdb.Client, org string, bucket string) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startTime := time.Now()
		err := handler(srv, ss)
		writePoint(startTime, in, org, bucket, info.FullMethod, err)
		return err
	}
}

func writePoint(startTime time.Time, in influxdb.Client, org, bucket, fullMethod string, err error) {
	if org != "" && bucket != "" {
		duration := time.Since(startTime)
		service := path.Dir(fullMethod)[1:]
		method := path.Base(fullMethod)

		writeAPI := in.WriteAPI(org, bucket)

		tags := map[string]string{
			"method":     method,
			"error_code": status.Code(err).String(),
		}
		fields := map[string]interface{}{
			"time_ms": duration.Milliseconds(),
		}

		p := influxdb.NewPoint(service, tags, fields, time.Now())
		writeAPI.WritePoint(p)
		writeAPI.Flush()
	}
}
