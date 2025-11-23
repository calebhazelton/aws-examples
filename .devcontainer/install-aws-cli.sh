cd /workspaces
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip -o awscliv2.zip
sudo ./aws/install
cd $CODESPACE_VSCODE_FOLDER

export AWS_ACCESS_KEY=$AWS_ACCESS_KEY_ENV
export AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY_ENV
export AWS_REGION=us-east-1