package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"net/http"
	"time"
)

const (
	ID              = "github"
	ClientIdKey     = "client_id"
	ClientSecretKey = "client_secret"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// User represents a GitHub user.
type User struct {
	Login                   *string    `json:"login,omitempty"`
	ID                      *int64     `json:"id,omitempty"`
	NodeID                  *string    `json:"node_id,omitempty"`
	AvatarURL               *string    `json:"avatar_url,omitempty"`
	HTMLURL                 *string    `json:"html_url,omitempty"`
	GravatarID              *string    `json:"gravatar_id,omitempty"`
	Name                    *string    `json:"name,omitempty"`
	Company                 *string    `json:"company,omitempty"`
	Blog                    *string    `json:"blog,omitempty"`
	Location                *string    `json:"location,omitempty"`
	Email                   *string    `json:"email,omitempty"`
	Hireable                *bool      `json:"hireable,omitempty"`
	Bio                     *string    `json:"bio,omitempty"`
	TwitterUsername         *string    `json:"twitter_username,omitempty"`
	PublicRepos             *int       `json:"public_repos,omitempty"`
	PublicGists             *int       `json:"public_gists,omitempty"`
	Followers               *int       `json:"followers,omitempty"`
	Following               *int       `json:"following,omitempty"`
	CreatedAt               *time.Time `json:"created_at,omitempty"`
	UpdatedAt               *time.Time `json:"updated_at,omitempty"`
	SuspendedAt             *time.Time `json:"suspended_at,omitempty"`
	Type                    *string    `json:"type,omitempty"`
	SiteAdmin               *bool      `json:"site_admin,omitempty"`
	TotalPrivateRepos       *int       `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos       *int       `json:"owned_private_repos,omitempty"`
	PrivateGists            *int       `json:"private_gists,omitempty"`
	DiskUsage               *int       `json:"disk_usage,omitempty"`
	Collaborators           *int       `json:"collaborators,omitempty"`
	TwoFactorAuthentication *bool      `json:"two_factor_authentication,omitempty"`

	// API URLs
	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
}

// Repository represents a GitHub repository.
type Repository struct {
	ID                  *int64           `json:"id,omitempty"`
	NodeID              *string          `json:"node_id,omitempty"`
	Owner               *User            `json:"owner,omitempty"`
	Name                *string          `json:"name,omitempty"`
	FullName            *string          `json:"full_name,omitempty"`
	Description         *string          `json:"description,omitempty"`
	Homepage            *string          `json:"homepage,omitempty"`
	DefaultBranch       *string          `json:"default_branch,omitempty"`
	MasterBranch        *string          `json:"master_branch,omitempty"`
	CreatedAt           *time.Time       `json:"created_at,omitempty"`
	PushedAt            *time.Time       `json:"pushed_at,omitempty"`
	UpdatedAt           *time.Time       `json:"updated_at,omitempty"`
	HTMLURL             *string          `json:"html_url,omitempty"`
	CloneURL            *string          `json:"clone_url,omitempty"`
	GitURL              *string          `json:"git_url,omitempty"`
	MirrorURL           *string          `json:"mirror_url,omitempty"`
	SSHURL              *string          `json:"ssh_url,omitempty"`
	SVNURL              *string          `json:"svn_url,omitempty"`
	Language            *string          `json:"language,omitempty"`
	Fork                *bool            `json:"fork,omitempty"`
	ForksCount          *int             `json:"forks_count,omitempty"`
	NetworkCount        *int             `json:"network_count,omitempty"`
	OpenIssuesCount     *int             `json:"open_issues_count,omitempty"`
	StargazersCount     *int             `json:"stargazers_count,omitempty"`
	SubscribersCount    *int             `json:"subscribers_count,omitempty"`
	WatchersCount       *int             `json:"watchers_count,omitempty"`
	Size                *int             `json:"size,omitempty"`
	AutoInit            *bool            `json:"auto_init,omitempty"`
	Parent              *Repository      `json:"parent,omitempty"`
	Source              *Repository      `json:"source,omitempty"`
	TemplateRepository  *Repository      `json:"template_repository,omitempty"`
	Permissions         *map[string]bool `json:"permissions,omitempty"`
	AllowRebaseMerge    *bool            `json:"allow_rebase_merge,omitempty"`
	AllowSquashMerge    *bool            `json:"allow_squash_merge,omitempty"`
	AllowMergeCommit    *bool            `json:"allow_merge_commit,omitempty"`
	DeleteBranchOnMerge *bool            `json:"delete_branch_on_merge,omitempty"`
	Topics              []string         `json:"topics,omitempty"`
	Archived            *bool            `json:"archived,omitempty"`
	Disabled            *bool            `json:"disabled,omitempty"`

	// Additional mutable fields when creating and editing a repository
	Private           *bool   `json:"private,omitempty"`
	HasIssues         *bool   `json:"has_issues,omitempty"`
	HasWiki           *bool   `json:"has_wiki,omitempty"`
	HasPages          *bool   `json:"has_pages,omitempty"`
	HasProjects       *bool   `json:"has_projects,omitempty"`
	HasDownloads      *bool   `json:"has_downloads,omitempty"`
	IsTemplate        *bool   `json:"is_template,omitempty"`
	LicenseTemplate   *string `json:"license_template,omitempty"`
	GitignoreTemplate *string `json:"gitignore_template,omitempty"`

	// API URLs
	URL              *string `json:"url,omitempty"`
	ArchiveURL       *string `json:"archive_url,omitempty"`
	AssigneesURL     *string `json:"assignees_url,omitempty"`
	BlobsURL         *string `json:"blobs_url,omitempty"`
	BranchesURL      *string `json:"branches_url,omitempty"`
	CollaboratorsURL *string `json:"collaborators_url,omitempty"`
	CommentsURL      *string `json:"comments_url,omitempty"`
	CommitsURL       *string `json:"commits_url,omitempty"`
	CompareURL       *string `json:"compare_url,omitempty"`
	ContentsURL      *string `json:"contents_url,omitempty"`
	ContributorsURL  *string `json:"contributors_url,omitempty"`
	DeploymentsURL   *string `json:"deployments_url,omitempty"`
	DownloadsURL     *string `json:"downloads_url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	ForksURL         *string `json:"forks_url,omitempty"`
	GitCommitsURL    *string `json:"git_commits_url,omitempty"`
	GitRefsURL       *string `json:"git_refs_url,omitempty"`
	GitTagsURL       *string `json:"git_tags_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssueCommentURL  *string `json:"issue_comment_url,omitempty"`
	IssueEventsURL   *string `json:"issue_events_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	KeysURL          *string `json:"keys_url,omitempty"`
	LabelsURL        *string `json:"labels_url,omitempty"`
	LanguagesURL     *string `json:"languages_url,omitempty"`
	MergesURL        *string `json:"merges_url,omitempty"`
	MilestonesURL    *string `json:"milestones_url,omitempty"`
	NotificationsURL *string `json:"notifications_url,omitempty"`
	PullsURL         *string `json:"pulls_url,omitempty"`
	ReleasesURL      *string `json:"releases_url,omitempty"`
	StargazersURL    *string `json:"stargazers_url,omitempty"`
	StatusesURL      *string `json:"statuses_url,omitempty"`
	SubscribersURL   *string `json:"subscribers_url,omitempty"`
	SubscriptionURL  *string `json:"subscription_url,omitempty"`
	TagsURL          *string `json:"tags_url,omitempty"`
	TreesURL         *string `json:"trees_url,omitempty"`
	TeamsURL         *string `json:"teams_url,omitempty"`
}

// Issue represents a GitHub issue on a repository.
type Issue struct {
	ID                *int64            `json:"id,omitempty"`
	Number            *int              `json:"number,omitempty"`
	State             *string           `json:"state,omitempty"`
	Locked            *bool             `json:"locked,omitempty"`
	Title             *string           `json:"title,omitempty"`
	Body              *string           `json:"body,omitempty"`
	AuthorAssociation *string           `json:"author_association,omitempty"`
	User              *User             `json:"user,omitempty"`
	Labels            []*Label          `json:"labels,omitempty"`
	Assignee          *User             `json:"assignee,omitempty"`
	Comments          *int              `json:"comments,omitempty"`
	ClosedAt          *time.Time        `json:"closed_at,omitempty"`
	CreatedAt         *time.Time        `json:"created_at,omitempty"`
	UpdatedAt         *time.Time        `json:"updated_at,omitempty"`
	ClosedBy          *User             `json:"closed_by,omitempty"`
	URL               *string           `json:"url,omitempty"`
	HTMLURL           *string           `json:"html_url,omitempty"`
	CommentsURL       *string           `json:"comments_url,omitempty"`
	EventsURL         *string           `json:"events_url,omitempty"`
	LabelsURL         *string           `json:"labels_url,omitempty"`
	RepositoryURL     *string           `json:"repository_url,omitempty"`
	Milestone         *Milestone        `json:"milestone,omitempty"`
	PullRequestLinks  *PullRequestLinks `json:"pull_request,omitempty"`
	Repository        *Repository       `json:"repository,omitempty"`
	Reactions         *Reactions        `json:"reactions,omitempty"`
	Assignees         []*User           `json:"assignees,omitempty"`
	NodeID            *string           `json:"node_id,omitempty"`
}

// Label represents a GitHub label on an Issue
type Label struct {
	ID          *int64  `json:"id,omitempty"`
	URL         *string `json:"url,omitempty"`
	Name        *string `json:"name,omitempty"`
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Default     *bool   `json:"default,omitempty"`
	NodeID      *string `json:"node_id,omitempty"`
}

// Milestone represents a GitHub repository milestone.
type Milestone struct {
	URL          *string    `json:"url,omitempty"`
	HTMLURL      *string    `json:"html_url,omitempty"`
	LabelsURL    *string    `json:"labels_url,omitempty"`
	ID           *int64     `json:"id,omitempty"`
	Number       *int       `json:"number,omitempty"`
	State        *string    `json:"state,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Creator      *User      `json:"creator,omitempty"`
	OpenIssues   *int       `json:"open_issues,omitempty"`
	ClosedIssues *int       `json:"closed_issues,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	ClosedAt     *time.Time `json:"closed_at,omitempty"`
	DueOn        *time.Time `json:"due_on,omitempty"`
	NodeID       *string    `json:"node_id,omitempty"`
}

// Reactions represents a summary of GitHub reactions.
type Reactions struct {
	TotalCount *int    `json:"total_count,omitempty"`
	PlusOne    *int    `json:"+1,omitempty"`
	MinusOne   *int    `json:"-1,omitempty"`
	Laugh      *int    `json:"laugh,omitempty"`
	Confused   *int    `json:"confused,omitempty"`
	Heart      *int    `json:"heart,omitempty"`
	Hooray     *int    `json:"hooray,omitempty"`
	Rocket     *int    `json:"rocket,omitempty"`
	Eyes       *int    `json:"eyes,omitempty"`
	URL        *string `json:"url,omitempty"`
}

// PullRequestLinks object is added to the Issue object when it's an issue included
// in the IssueCommentEvent webhook payload, if the webhook is fired by a comment on a PR.
type PullRequestLinks struct {
	URL      *string `json:"url,omitempty"`
	HTMLURL  *string `json:"html_url,omitempty"`
	DiffURL  *string `json:"diff_url,omitempty"`
	PatchURL *string `json:"patch_url,omitempty"`
}

// Project represents a GitHub Project.
type Project struct {
	ID         *int64  `json:"id,omitempty"`
	URL        *string `json:"url,omitempty"`
	HTMLURL    *string `json:"html_url,omitempty"`
	ColumnsURL *string `json:"columns_url,omitempty"`
	OwnerURL   *string `json:"owner_url,omitempty"`
	Name       *string `json:"name,omitempty"`
	Body       *string `json:"body,omitempty"`
	Number     *int    `json:"number,omitempty"`
	State      *string `json:"state,omitempty"`
	NodeID     *string `json:"node_id,omitempty"`

	// The User object that generated the project.
	Creator *User `json:"creator,omitempty"`
}

// ProjectColumn represents a column of a GitHub Project.
type ProjectColumn struct {
	ID         *int64  `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	URL        *string `json:"url,omitempty"`
	ProjectURL *string `json:"project_url,omitempty"`
	CardsURL   *string `json:"cards_url,omitempty"`
	NodeID     *string `json:"node_id,omitempty"`
}

// ProjectCard represents a card in a column of a GitHub Project.
type ProjectCard struct {
	URL        *string `json:"url,omitempty"`
	ColumnURL  *string `json:"column_url,omitempty"`
	ContentURL *string `json:"content_url,omitempty"`
	ID         *int64  `json:"id,omitempty"`
	Note       *string `json:"note,omitempty"`
	Creator    *User   `json:"creator,omitempty"`
	NodeID     *string `json:"node_id,omitempty"`
	Archived   *bool   `json:"archived,omitempty"`

	// The following fields are only populated by Webhook events.
	ColumnID *int64 `json:"column_id,omitempty"`

	// The following fields are only populated by Events API.
	ProjectID          *int64  `json:"project_id,omitempty"`
	ProjectURL         *string `json:"project_url,omitempty"`
	ColumnName         *string `json:"column_name,omitempty"`
	PreviousColumnName *string `json:"previous_column_name,omitempty"` // Populated in "moved_columns_in_project" event deliveries.
}

type Github struct {
	c            *resty.Client
	clientId     string
	clientSecret string
	redirectURI  string
	accessToken  string
}

func NewGithub(clientId, clientSecret, redirectURI, accessToken string) *Github {
	v := &Github{clientId: clientId, clientSecret: clientSecret, redirectURI: redirectURI, accessToken: accessToken}

	v.c = resty.New()
	v.c.SetHostURL("https://api.github.com")
	v.c.SetTimeout(time.Minute)

	return v
}

func (v *Github) AuthorizeURL() string {
	return fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=repo", v.clientId, v.redirectURI)
}

func (v *Github) GetAccessToken(code string) (interface{}, error) {
	resp, err := v.c.R().
		SetResult(&TokenResponse{}).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetBody(map[string]interface{}{
			"client_id":     v.clientId,
			"client_secret": v.clientSecret,
			"code":          code,
		}).
		Post("https://github.com/login/oauth/access_token")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		result := resp.Result().(*TokenResponse)
		v.accessToken = result.AccessToken
		return result, nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}

func (v *Github) Redirect(c *fiber.Ctx, mid pb.MiddleClient) error {
	reply, err := mid.GetCredential(context.Background(), &pb.CredentialRequest{Type: ID})
	if err != nil {
		return err
	}
	clientId := ""
	for _, item := range reply.GetContent() {
		if item.Key == ClientIdKey {
			clientId = item.Value
		}
	}
	v.clientId = clientId

	appRedirectURI := v.AuthorizeURL()
	return c.Redirect(appRedirectURI, http.StatusFound)
}

func (v *Github) StoreAccessToken(c *fiber.Ctx, mid pb.MiddleClient) error {
	code := c.FormValue("code")
	reply, err := mid.GetCredential(context.Background(), &pb.CredentialRequest{Type: ID})
	if err != nil {
		return err
	}
	clientId := ""
	clientSecret := ""
	for _, item := range reply.GetContent() {
		if item.Key == ClientIdKey {
			clientId = item.Value
		}
		if item.Key == ClientSecretKey {
			clientSecret = item.Value
		}
	}
	v.clientId = clientId
	v.clientSecret = clientSecret

	tokenResp, err := v.GetAccessToken(code)
	if err != nil {
		return err
	}

	extra, err := json.Marshal(&tokenResp)
	if err != nil {
		return err
	}
	appReply, err := mid.StoreAppOAuth(context.Background(), &pb.AppRequest{
		Name:  ID,
		Type:  ID,
		Token: v.accessToken,
		Extra: utils.ByteToString(extra),
	})
	if err != nil {
		return err
	}
	if appReply.GetState() {
		return nil
	}
	return errors.New("error")
}

func (v *Github) GetUser() (*User, error) {
	resp, err := v.c.R().
		SetResult(&User{}).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetHeader("Authorization", fmt.Sprintf("token %s", v.accessToken)).
		Get("/user")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Result().(*User), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}

func (v *Github) GetStarred(username string) (*[]Repository, error) {
	resp, err := v.c.R().
		SetResult(&[]Repository{}).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetHeader("Authorization", fmt.Sprintf("token %s", v.accessToken)).
		Get(fmt.Sprintf("/users/%s/starred", username))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Result().(*[]Repository), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}

func (v *Github) GetFollowers() (*[]User, error) {
	resp, err := v.c.R().
		SetResult(&[]User{}).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetHeader("Authorization", fmt.Sprintf("token %s", v.accessToken)).
		Get("/user/followers")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Result().(*[]User), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}

func (v *Github) CreateIssue(owner, repo string, issue Issue) (*Issue, error) {
	resp, err := v.c.R().
		SetResult(&Issue{}).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetHeader("Authorization", fmt.Sprintf("token %s", v.accessToken)).
		SetBody(issue).
		Post(fmt.Sprintf("/repos/%s/%s/issues", owner, repo))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusCreated {
		return resp.Result().(*Issue), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}

func (v *Github) GetUserProjects(username string) (*[]Project, error) {
	resp, err := v.c.R().
		SetResult(&[]Project{}).
		SetHeader("Accept", "application/vnd.github.inertia-preview+json").
		SetHeader("Authorization", fmt.Sprintf("token %s", v.accessToken)).
		Get(fmt.Sprintf("/users/%s/projects", username))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Result().(*[]Project), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}

func (v *Github) GetProjectColumns(projectID int64) (*[]ProjectColumn, error) {
	resp, err := v.c.R().
		SetResult(&[]ProjectColumn{}).
		SetHeader("Accept", "application/vnd.github.inertia-preview+json").
		SetHeader("Authorization", fmt.Sprintf("token %s", v.accessToken)).
		Get(fmt.Sprintf("/projects/%d/columns", projectID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return resp.Result().(*[]ProjectColumn), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}

func (v *Github) CreateCard(columnID int64, card ProjectCard) (*ProjectCard, error) {
	resp, err := v.c.R().
		SetResult(&ProjectCard{}).
		SetHeader("Accept", "application/vnd.github.inertia-preview+json").
		SetHeader("Authorization", fmt.Sprintf("token %s", v.accessToken)).
		SetBody(card).
		Post(fmt.Sprintf("/projects/columns/%d/cards", columnID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusCreated {
		return resp.Result().(*ProjectCard), nil
	} else {
		return nil, fmt.Errorf("%d, %s (%s)", resp.StatusCode(), resp.Header().Get("X-Error-Code"), resp.Header().Get("X-Error"))
	}
}
