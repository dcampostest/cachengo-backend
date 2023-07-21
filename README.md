# cachengo-backend

Run Debuguin code in local 
CompileDaemon -command="./golan-rest-simple"

# cachengo-backend run
nohup go run main.go &

ps -ef | grep go
kill PID



To run the Go App as a service you need to create a new service unit file.

However, the App needs to know where Go is installed. The easiest way to lookup that location is by running this command:

which go
which gives you an output like this:

/usr/local/go/bin/go
With this piece of information, you can create the systemd service file. Create a file named providus-app.service in the /etc/systemd/system/ using the command below:

sudo touch /etc/systemd/system/cachengo-backend.service
Next open the newly created file:

sudo nano /etc/systemd/system/cachengo-backend.service
Paste the following configuration into your service file:

[Unit]
Description=Cachengo Backend service
After=network.target

[Service]
Type=forking

User=deploy
Group=deploy

ExecStart=/usr/local/go/bin/go run main.go
WorkingDirectory=/root/cachengo-backend

Restart=always
RestartSec=10
KillSignal=SIGINT

SyslogIdentifier=cachengo-backend-service
PrivateTmp=true

[Install]
WantedBy=multi-user.target
When you are finished, save and close the file.

Next, reload the systemd daemon so that it knows about our service file:

sudo systemctl daemon-reload
Start the Providus App service by typing:

sudo systemctl restart cachengo-backend
Double-check that it started without errors by typing:

sudo systemctl status cachengo-backend
And then enable the Providus App service file so that Providus App automatically starts at boot, that is, it can start on its own whenever the server restarts:

sudo systemctl enable cachengo-backend
This creates a multi-user.target symlink in /etc/systemd/system/multi-user.target.wants/providus-app.service for the /etc/systemd/system/providus-app.service file that you created.

To check logs:

sudo journalctl -u cachengo-backend