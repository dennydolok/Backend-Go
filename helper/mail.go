package helper

import "github.com/mailjet/mailjet-apiv3-go"

func SendMail(code, email, name, context string) error {
	publicKey := "b5cd4a33c4ea6788fbdc347067b3a35b"
	secretKey := "4fbc6b06490243458fe04b3182f3f818"
	mj := mailjet.NewMailjetClient(publicKey, secretKey)
	messageInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "bearuang0816@gmail.com",
				Name:  "WallE",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  name,
				},
			},
			Subject:  "Verifikasi Kode",
			TextPart: "Berikut verifikasi kode anda untuk " + context + "! \n " + code,
			HTMLPart: "<h3>Berikut verifikasi kode anda untuk " + context + "!</h3> <br /><center><strong>" + code + "</strong></center>",
		},
	}
	messages := mailjet.MessagesV31{Info: messageInfo}
	_, err := mj.SendMailV31(&messages)
	if err != nil {
		return err
	}
	return nil
}
