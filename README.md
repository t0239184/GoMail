# GoMail
Send mail with Golang

## 實作過程中發現的小細節
看Go的smtp文件中有提到如果使用msg這個欄位, 需要加上"From", "To", "Subject"這些欄位在Msg中
> The msg headers should usually include fields such as "From", "To", "Subject", and "Cc". Sending "Bcc" messages is accomplished by including an email address in the to parameter but not including it in the msg headers.
> https://pkg.go.dev/net/smtp


然後發現在"From", "To"這兩個欄位的後面需要加上\n才可以, 如果加上\n\n會變成之後的都會被視為Body, 所以就是在Body前的那個欄位要使用\n\n