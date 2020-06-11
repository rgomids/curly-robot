#! /bin/bash
sudo rm -rf /etc/$TURBO_PARAKEET_DEFAULT_NAME;
sudo rm -rf /opt/$TURBO_PARAKEET_DEFAULT_NAME;
sudo rm -rf /var/log/$TURBO_PARAKEET_DEFAULT_NAME;
sudo mkdir /etc/$TURBO_PARAKEET_DEFAULT_NAME;
ssh-keygen -t rsa -b 4096 -N "" -C "$TURBO_PARAKEET_EMAIL" -f ./static/$TURBO_PARAKEET_DEFAULT_NAME;
sudo cp ./static/configs.json /etc/$TURBO_PARAKEET_DEFAULT_NAME/;
sudo mkdir /opt/$TURBO_PARAKEET_DEFAULT_NAME;
sudo mkdir /opt/$TURBO_PARAKEET_DEFAULT_NAME/.key;
sudo cp -r ./bin /opt/$TURBO_PARAKEET_DEFAULT_NAME;
sudo cp ./static/$TURBO_PARAKEET_DEFAULT_NAME /opt/$TURBO_PARAKEET_DEFAULT_NAME/.key;
sudo chown $USER:$USER /opt/$TURBO_PARAKEET_DEFAULT_NAME/.key/$TURBO_PARAKEET_DEFAULT_NAME;
sudo touch /var/log/$TURBO_PARAKEET_DEFAULT_NAME;
sudo chown $USER:$USER /var/log/$TURBO_PARAKEET_DEFAULT_NAME;
sudo chmod 0755 /var/log/$TURBO_PARAKEET_DEFAULT_NAME;