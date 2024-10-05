import subprocess
import os

def generate_keys():
    private_key = subprocess.check_output(["wg", "genkey"]).decode("utf-8").strip()
    public_key = subprocess.check_output(["wg", "pubkey"], input=private_key.encode()).decode("utf-8").strip()
    return public_key, private_key

def generate_config(client):
    config = f"""
[Interface]
PrivateKey = {client.private_key}
Address = {client.ip_address}/24
DNS = 8.8.8.8

[Peer]
PublicKey = {os.environ.get('SERVER_PUBLIC_KEY')}
AllowedIPs = 0.0.0.0/0
Endpoint = {os.environ.get('SERVER_ENDPOINT')}
PersistentKeepalive = 25
"""
    return config