package gmail

import (
	"fmt"
	"github.com/mchirico/go-gmail/mail/creds"
	"google.golang.org/api/gmail/v1"
)


func Labels() (map[string]string, error) {

	srv := creds.NewGmailSrv()
	user := "me"
	m := map[string]string{}
	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		return m, err
	}
	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return m, err
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		m[l.Name] = l.Id
		fmt.Printf("- %s, %s\n", l.Name, l.Id)
	}
	return m, nil
}

func Delete(labelID string) (int64, error) {
	var count int64
	srv := creds.NewGmailSrv()
	nsrv := gmail.NewUsersService(srv)
	msg, err := nsrv.Messages.List("me").LabelIds(labelID).Do()
	if err != nil {
		return count, err
	}
	for _, v := range msg.Messages {
		count += 1
		nsrv.Messages.Delete("me", v.Id).Do()
	}
	return count, err
}

func BigDelete() {
	Delete("INBOX")
	Delete("TRASH")
	Delete("SPAM")
}