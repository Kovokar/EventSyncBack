erDiagram
    USUARIO ||--o{ PARTIDA : "cria"
    USUARIO ||--o{ PARTICIPACAO_PARTIDA : "participa"
    USUARIO ||--o{ RESERVA : "realiza"
    USUARIO ||--o{ AVALIACAO : "avalia"
    USUARIO ||--o{ AVALIACAO : "recebe"
    USUARIO ||--o{ MEMBRO_TIME : "pertence"
    USUARIO ||--o{ TIME : "cria"

    ESTABELECIMENTO ||--o{ ARENA : "possui"
    
    ARENA ||--o{ RESERVA : "tem"
    ARENA ||--o{ HORARIO_DISPONIVEL : "possui"
    ARENA ||--o{ PARTIDA : "sedia"
    
    PARTIDA ||--o{ PARTICIPACAO_PARTIDA : "tem"
    
    TIME ||--o{ MEMBRO_TIME : "tem"
    TIME ||--o{ TREINAMENTO : "realiza"
    TIME ||--o{ COMPETICAO : "participa"
    
    ARENA ||--o{ TREINAMENTO : "sedia"

    USUARIO {
        int id PK
        string nome
        int idade
        string email UK
        string senha
        string foto
        decimal latitude
        decimal longitude
        datetime created_at
        datetime updated_at
    }
    
    ESTABELECIMENTO {
        int id PK
        string nome
        string endereco
        decimal latitude
        decimal longitude
        string telefone
        string email
        string cnpj UK
        datetime created_at
        datetime updated_at
    }
    
    ARENA {
        int id PK
        int estabelecimento_id FK
        string nome
        int capacidade_jogadores
        decimal preco_hora
        string descricao
        boolean ativa
        datetime created_at
        datetime updated_at
    }
    
    HORARIO_DISPONIVEL {
        int id PK
        int arena_id FK
        date data
        time hora_inicio
        time hora_fim
        boolean disponivel
        datetime created_at
    }
    
    RESERVA {
        int id PK
        int usuario_id FK
        int arena_id FK
        date data
        time hora_inicio
        time hora_fim
        decimal valor_total
        string status
        datetime created_at
        datetime updated_at
    }
    
    PARTIDA {
        int id PK
        int criador_id FK
        int arena_id FK
        int reserva_id FK
        string tipo
        date data
        time hora_inicio
        time hora_fim
        int capacidade_maxima
        string codigo_convite
        string status
        datetime created_at
        datetime updated_at
    }
    
    PARTICIPACAO_PARTIDA {
        int id PK
        int partida_id FK
        int usuario_id FK
        string status
        string time
        boolean vencedor
        datetime data_entrada
        datetime created_at
    }
    
    AVALIACAO {
        int id PK
        int avaliador_id FK
        int avaliado_id FK
        int partida_id FK
        int nota
        string comentario
        datetime created_at
    }
    
    TIME {
        int id PK
        int criador_id FK
        string nome
        string descricao
        string logo
        datetime created_at
        datetime updated_at
    }
    
    MEMBRO_TIME {
        int id PK
        int time_id FK
        int usuario_id FK
        string papel
        datetime data_entrada
        boolean ativo
        datetime created_at
    }
    
    TREINAMENTO {
        int id PK
        int time_id FK
        int arena_id FK
        date data
        time hora_inicio
        time hora_fim
        string descricao
        datetime created_at
    }
    
    COMPETICAO {
        int id PK
        int time_id FK
        string nome
        string descricao
        date data_inicio
        date data_fim
        string resultado
        datetime created_at
    }
