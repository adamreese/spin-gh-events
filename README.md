# Example Spin app for GitHub webhooks

### Run the app on local machine

```
$ spin build --up --key-value='webhook_secret=YOURSECRETKEY'
```

### Running the app on Fermyon Cloud

First create a [webhook on GitHub](https://docs.github.com/en/webhooks-and-events/webhooks/creating-webhooks) and copy the secret key.  Use Spin to build and
deploy.

```
$ spin build
$ spin deploy --key-value='webhook_secret=YOURSECRETKEY'
```

You now have an application accepting webhooks from your GitHub repository!

## Testing mock payloads

Use the test script to send mock payloads to your application.

```
$ export ENDPOINT=myapp.fermyon.app
$ export SECRET_KEY=YOURSECRETKEY
$ ./test.sh
```

Get the logs on [https://cloud.fermyon.com](https://cloud.fermyon.com).
