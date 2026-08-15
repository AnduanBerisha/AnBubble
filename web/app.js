'use strict';

const messagesEl = document.getElementById('messages');
const form = document.getElementById('composer');
const input = document.getElementById('text-input');
const seen = new Set();

let nickname = localStorage.getItem('anbubble_nick');
while (!nickname) {
  nickname = (prompt('What is your name ? :') || '').trim().slice(0, 25);
}
localStorage.setItem('anbubble_nick', nickname);

async function loadMessages() {
  try {
    const res = await fetch('/messages');
    const data = await res.json();
    for (const msg of data) {
      if (seen.has(msg.id)) continue;
      seen.add(msg.id);
      addBubble(msg);
    }
  } catch (err) {
    console.error('Cannot load', err);
  }
}

function addBubble(msg) {
  const bubble = document.createElement('div');
  bubble.className = 'bubble' + (msg.sender === nickname ? ' mine' : '');

  const sender = document.createElement('div');
  sender.className = 'sender';
  sender.textContent = msg.sender;

  const text = document.createElement('div');
  text.textContent = msg.text;

  bubble.append(sender, text);
  messagesEl.appendChild(bubble);
  messagesEl.scrollTop = messagesEl.scrollHeight;
}

form.addEventListener('submit', async (e) => {
  e.preventDefault();
  const text = input.value.trim();
  if (!text) return;

  try {
    await fetch('/send', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ sender: nickname, text }),
    });
    input.value = '';
    loadMessages();
  } catch (err) {
    console.error('Cannot send', err);
  }
});

loadMessages();
setInterval(loadMessages, 2000);
