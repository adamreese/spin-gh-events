# Example Spin app for GitHub webhooks

## Running the app on Fermyon Cloud

First create a [webhook on GitHub](https://docs.github.com/en/webhooks-and-events/webhooks/creating-webhooks) and copy the secret key.  Use Spin to build and
deploy.

```
$ spin build
$ spin deploy --key-value='webhookSecretKey=YOURSECRETKEY'
```
You now have an application accepting webhooks from your GitHub repository!

__Testing without GitHub__

Use the test script to send mock payloads to your application.

```
export ENDPOINT=myapp.fermyon.app
./test.sh
```

Get the logs on [https://cloud.fermyon.com](https://cloud.fermyon.com).
