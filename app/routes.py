from flask import Blueprint, jsonify, request, render_template, redirect, url_for
from app import db
from app.models import WireGuardClient
from app.wireguard import generate_keys, generate_config

bp = Blueprint('main', __name__)

@bp.route('/')
def index():
    clients = WireGuardClient.query.all()
    return render_template('index.html', clients=clients)

@bp.route('/add_client', methods=['GET', 'POST'])
def add_client():
    if request.method == 'POST':
        name = request.form['name']
        public_key, private_key = generate_keys()
        ip_address = f'10.0.0.{WireGuardClient.query.count() + 2}'
        
        new_client = WireGuardClient(name=name, public_key=public_key, private_key=private_key, ip_address=ip_address)
        db.session.add(new_client)
        db.session.commit()
        
        return redirect(url_for('main.index'))
    
    return render_template('add_client.html')

@bp.route('/remove_client/<int:id>', methods=['POST'])
def remove_client(id):
    client = WireGuardClient.query.get_or_404(id)
    db.session.delete(client)
    db.session.commit()
    return redirect(url_for('main.index'))

@bp.route('/api/clients', methods=['GET'])
def get_clients():
    clients = WireGuardClient.query.all()
    return jsonify([{'name': c.name, 'public_key': c.public_key, 'ip_address': c.ip_address} for c in clients])

@bp.route('/api/client/<int:id>/config', methods=['GET'])
def get_client_config(id):
    client = WireGuardClient.query.get_or_404(id)
    config = generate_config(client)
    return jsonify({'config': config})