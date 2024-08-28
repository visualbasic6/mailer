# simple mailgun mailer

~

`snap install go --classic`

`git clone https://github.com/visualbasic6/mailer.git`

`touch subjects.txt && touch body.html && touch recipients.txt`

`nano` whatever`.txt`

`/subjects.txt` = 1 line, the subject

`/body.html` = the email body

`/recipients.txt` = email addresses, add 'email' to the top of the file

`go run mailer.go`

throttling can be modified in the code by beginners. happy blasting!
