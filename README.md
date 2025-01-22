[//]: # (Copyright [c] 2025 sigma-firma)

[//]: # (Permission is hereby granted, free of charge, to any person obtaining a copy)
[//]: # (of this software and associated documentation files [the "Software"], to deal)
[//]: # (in the Software without restriction, including without limitation the rights)
[//]: # (to use, copy, modify, merge, publish, distribute, sublicense, and/or sell)
[//]: # (copies of the Software, and to permit persons to whom the Software is)
[//]: # (furnished to do so, subject to the following conditions:)

[//]: # (The above copyright notice and this permission notice shall be included in all)
[//]: # (copies or substantial portions of the Software.)

[//]: # (THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR)
[//]: # (IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,)
[//]: # (FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE)
[//]: # (AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER)
[//]: # (LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,)
[//]: # (OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE)
[//]: # (SOFTWARE.)

# gmailAPI

A go package for managing gmail Tokens/credentials. 

All credit to the google gmail API devs who provided an example here: 
https://developers.google.com/gmail/api/quickstart/go

This package was adapted from that example and built in a way that allows the 
user to import it into any project that uses the gmail API. This package first 
checks for credentials in a users home directory, if they aren't found it 
attempts to create the credentials by prompting the user to visit a link and 
paste the authentication token found at that link into the program. This only 
needs to be done once. The next version of gmailAPI will include support for 
multiple users. 

NOTE: Often, the web page google uses to provide credentials appears broken, 
and unable to load. Examine the web address carefully, the token can be copied
from there.

# USE:

	package gmailAPI

	import (
			"context"
			"fmt"

			"github.com/sigma-firma/gmailAPI"
			gmail "google.golang.org/api/gmail/v1"
	)

	func main() {
			// Connect to the Gmail API service. Here, we use a context and provide a
			// scope. The scope is used by the Gamil API to determine your privilege
			// level. gmailAPI.ConnectToService is a variadic function and thus can be
			// passed any number of scopes. For more information on scopes see:
			// https://developers.google.com/gmail/api/auth/scopes
			ctx := context.Background()
			srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)

			// Get a list of your unread messages
			inbox, err := srv.Users.Messages.List("me").Q("in:UNREAD").Do()
			if err != nil {
;;					fmt.Println(err)
			}

			for _, message := range inbox.Messages {
					// To get the message content, you must make a second call
					// to the gmail API for each individual ID.
					msg, _ := srv.Users.Messages.Get("me", message.Id).Do()
					fmt.Println(msg.Snippet)
			}
	}


## CREDENTIALS:

For gmailAPI to work you must have a gmail account and a file named 
"client_secret.json" containing your authorization info in the root directory 
of your project. To obtain credentials please see step one of this guide: 
https://developers.google.com/gmail/api/quickstart/go

#### Turning on the gmail API

  - Use this wizard (https://console.developers.google.com/start/api?id=gmail) 
  to create or select a project in the Google Developers Console and 
  automatically turn on the API. Click Continue, 
  then Go to credentials.
 
  - On the Add credentials to your project page, click the Cancel button.
 
  - At the top of the page, select the OAuth consent screen tab. Select an 
  Email address, enter a Product name if not already set, and click the Save 
  button.
 
  - Select the Credentials tab, click the Create credentials button and select
  OAuth client ID.
 
  - Select the application type Other, enter the name "Gmail API Quickstart", 
  and click the Create button.
 
  - Click OK to dismiss the resulting dialog.
 
  - Click the file_download (Download JSON) button to the right of the client 
  ID.
 
  - Move this file to your working directory and rename it client_secret.json.


Also see our other package, [inboxer](https://gitlab.com/sigma-firma/inboxer), 
which makes performing basic actions on your inbox much more straight-forward. 

