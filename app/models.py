from app import db

class WireGuardClient(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(100), unique=True, nullable=False)
    public_key = db.Column(db.String(44), unique=True, nullable=False)
    private_key = db.Column(db.String(44), unique=True, nullable=False)
    ip_address = db.Column(db.String(15), unique=True, nullable=False)

    def __repr__(self):
        return f'<WireGuardClient {self.name}>'