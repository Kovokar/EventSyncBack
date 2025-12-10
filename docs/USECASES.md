# Diagrama de Caso de Uso - Rede Social para Formação de Times de Vôlei

## Atores
- **Usuário**
- **Estabelecimento**
- **Sistema**

---

## Casos de Uso

### 1. **Cadastro e Login**
#### Ator: **Usuário**
- O usuário realiza o cadastro fornecendo nome, idade, e e-mail.
- O usuário compartilha sua localização (latitude e longitude).
- O usuário faz login para acessar a plataforma.

---

### 2. **Busca e Reserva de Arenas**
#### Ator: **Usuário**
- O usuário pode buscar estabelecimentos e arenas com base na sua localização.
- O usuário pode visualizar a disponibilidade das arenas.
- O usuário pode reservar uma arena, conforme a disponibilidade e capacidade.

#### Ator: **Estabelecimento**
- O estabelecimento define a capacidade das arenas e a disponibilidade de horários.

---

### 3. **Criação de Partidas**
#### Ator: **Usuário**
- O usuário cria uma partida (pública ou privada).
  - **Pública:** O criador aguarda solicitações de entrada e decide quem pode participar.
  - **Privada:** O criador envia um código de convite para que outros usuários participem.

#### Ator: **Sistema**
- O sistema valida o número de participantes conforme a capacidade da arena.
- O sistema confirma a reserva da arena.

---

### 4. **Perfil de Usuário e Histórico**
#### Ator: **Usuário**
- O usuário pode criar um perfil pessoal com foto, informações básicas e histórico de partidas.
- O usuário pode visualizar seu histórico de jogos (vitórias, derrotas, e participação).
- O usuário pode ver seu ranking com base no desempenho (vitórias, participação, etc.).
- O usuário pode avaliar outros jogadores após a partida.

---

### 5. **Criação e Gerenciamento de Times**
#### Ator: **Usuário**
- O usuário pode criar um time de vôlei ou se juntar a um time existente.
- O usuário pode convidar outros jogadores para o time.
- O usuário pode visualizar o calendário de competições do time.
- O usuário pode participar de sessões de treinamento organizadas pelo time.

#### Ator: **Estabelecimento**
- O estabelecimento pode fornecer a opção de reservar arenas para treinos de times.

#### Ator: **Sistema**
- O sistema registra o time, seus membros, e o histórico de jogos.
- O sistema mantém um calendário de jogos e treinos para cada time.

---

### 6. **Avaliação de Jogadores**
#### Ator: **Usuário**
- Após cada partida, o usuário pode avaliar outros jogadores.
- O usuário pode ser avaliado pelos outros jogadores após cada partida.

#### Ator: **Sistema**
- O sistema calcula o ranking dos jogadores com base nas avaliações e desempenho.

---

## Diagrama de Casos de Uso

```plaintext
+---------------------------------------------------------------+
|                     Rede Social de Vôlei                      |
+---------------------------------------------------------------+
|                               Sistema                          |
+---------------------------------------------------------------+
| 1. Cadastro e Login                    2. Busca e Reserva de Arenas |
|   (Usuário)                            (Usuário, Estabelecimento)   |
|                                         +-------------------------+|
| 3. Criação de Partidas                   4. Perfil e Histórico     |
|   (Usuário, Sistema)                    (Usuário)                  |
|                                         +-------------------------+|
| 5. Criação e Gerenciamento de Times    6. Avaliação de Jogadores  |
|   (Usuário, Estabelecimento, Sistema)   (Usuário, Sistema)         |
+---------------------------------------------------------------+
