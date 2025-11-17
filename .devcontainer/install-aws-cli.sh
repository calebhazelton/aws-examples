cd /workspaces
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
cd $CODESPACE_VSCODE_FOLDER

export $AWS_ACCESS_KEY
export $AWS_SECRET_ACCESS_KEY
export $AWS_REGION