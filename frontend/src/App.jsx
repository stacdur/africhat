import React from 'react';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>🌍 AfriChat</h1>
        <p>Application de messagerie pour l'Afrique</p>
        <div className="features">
          <div className="feature">📞 Appels Audio/Vidéo</div>
          <div className="feature">💰 Mobile Money</div>
          <div className="feature">🌐 5 Langues</div>
          <div className="feature">📶 Mode Offline</div>
        </div>
        <p className="author">Par stacdur</p>
      </header>
    </div>
  );
}

export default App;
