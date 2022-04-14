package sendgrid

import (
	"github.com/porter-dev/porter/internal/notifier"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type UserNotifier struct {
	client *Client
}

type Client struct {
	APIKey                  string
	PWResetTemplateID       string
	PWGHTemplateID          string
	VerifyEmailTemplateID   string
	ProjectInviteTemplateID string
	SenderEmail             string
}

func NewUserNotifier(client *Client) notifier.UserNotifier {
	return &UserNotifier{client}
}

func (s *UserNotifier) SendPasswordResetEmail(opts *notifier.SendPasswordResetEmailOpts) error {
	request := sendgrid.GetRequest(s.client.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	sgMail := &mail.SGMailV3{
		Personalizations: []*mail.Personalization{
			{
				To: []*mail.Email{
					{
						Address: opts.Email,
					},
				},
				DynamicTemplateData: map[string]interface{}{
					"url":   opts.URL,
					"email": opts.Email,
				},
			},
		},
		From: &mail.Email{
			Address: s.client.SenderEmail,
			Name:    "Porter",
		},
		TemplateID: s.client.PWResetTemplateID,
	}

	request.Body = mail.GetRequestBody(sgMail)

	_, err := sendgrid.API(request)

	return err
}

func (s *UserNotifier) SendGithubRelinkEmail(opts *notifier.SendGithubRelinkEmailOpts) error {
	request := sendgrid.GetRequest(s.client.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	sgMail := &mail.SGMailV3{
		Personalizations: []*mail.Personalization{
			{
				To: []*mail.Email{
					{
						Address: opts.Email,
					},
				},
				DynamicTemplateData: map[string]interface{}{
					"url":   opts.URL,
					"email": opts.Email,
				},
			},
		},
		From: &mail.Email{
			Address: s.client.SenderEmail,
			Name:    "Porter",
		},
		TemplateID: s.client.PWGHTemplateID,
	}

	request.Body = mail.GetRequestBody(sgMail)

	_, err := sendgrid.API(request)

	return err
}

func (s *UserNotifier) SendEmailVerification(opts *notifier.SendEmailVerificationOpts) error {
	request := sendgrid.GetRequest(s.client.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	sgMail := &mail.SGMailV3{
		Personalizations: []*mail.Personalization{
			{
				To: []*mail.Email{
					{
						Address: opts.Email,
					},
				},
				DynamicTemplateData: map[string]interface{}{
					"url":   opts.URL,
					"email": opts.Email,
				},
			},
		},
		From: &mail.Email{
			Address: s.client.SenderEmail,
			Name:    "Porter",
		},
		TemplateID: s.client.VerifyEmailTemplateID,
	}

	request.Body = mail.GetRequestBody(sgMail)

	_, err := sendgrid.API(request)

	return err
}

func (s *UserNotifier) SendProjectInviteEmail(opts *notifier.SendProjectInviteEmailOpts) error {
	request := sendgrid.GetRequest(s.client.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	sgMail := &mail.SGMailV3{
		Personalizations: []*mail.Personalization{
			{
				To: []*mail.Email{
					{
						Address: opts.InviteeEmail,
					},
				},
				DynamicTemplateData: map[string]interface{}{
					"url":          opts.URL,
					"sender_email": opts.ProjectOwnerEmail,
					"project":      opts.Project,
				},
			},
		},
		From: &mail.Email{
			Address: s.client.SenderEmail,
			Name:    "Porter",
		},
		TemplateID: s.client.ProjectInviteTemplateID,
	}

	request.Body = mail.GetRequestBody(sgMail)

	_, err := sendgrid.API(request)

	return err
}
