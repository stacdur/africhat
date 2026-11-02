# 🌍 AfriChat - Application de Messagerie pour l'Afrique

![AfriChat](https://img.shields.io/badge/AfriChat-Pour%20l'Afrique-22c55e?style=for-the-badge&logo=whatsapp&logoColor=white)

[![GitHub Stars](https://img.shields.io/github/stars/stacdur/africhat?style=social)](https://github.com/stacdur/africhat)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

> Application de messagerie complète avec appels audio/vidéo sécurisés, Mobile Money intégré et support multilingue

## ✨ Fonctionnalités

- 💬 **Messagerie instantanée** - Texte, images, vidéos, audio
- 📞 **Appels audio/vidéo** - WebRTC avec chiffrement E2E
- 💰 **Mobile Money** - MTN, Orange Money, M-Pesa, Airtel, Moov
- 🌐 **5 langues** - Français, English, Swahili, Wolof, Hausa
- 📶 **Mode ultra-léger** - Optimisé pour 2G/3G
- 💾 **Mode offline** - Synchronisation automatique
- 🔒 **Sécurisé** - JWT, HTTPS, chiffrement

## 🚀 Démarrage Rapide

### Option 1 : Docker (Recommandé)
```bash
git clone https://github.com/stacdur/africhat.git
cd africhat
docker-compose up -d
```

### Option 2 : Manuel

**Backend:**
```bash
cd backend
go mod download
cp .env.example .env
go run cmd/server/main.go
```

**Frontend:**
```bash
cd frontend
npm install
cp .env.example .env
npm start
```

## 📖 Documentation

- [Installation](docs/INSTALLATION.md)
- [Déploiement](docs/DEPLOYMENT.md)
- [API Documentation](docs/API.md)
- [Contribution](CONTRIBUTING.md)

## 👤 Auteur

**stacdur**
- GitHub: [@stacdur](https://github.com/stacdur)
- Email: stacdur@gmail.com

## 📄 License

MIT © 2025 stacdur

---

<div align="center">

**Fait avec ❤️ pour l'Afrique**

[⭐ Star ce projet](https://github.com/stacdur/africhat) si vous l'aimez !

</div>
