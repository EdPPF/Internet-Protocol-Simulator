## Como rodar o projeto

Clone o repositório:

```bash
git clone https://github.com/EdPPF/Internet-Protocol-Simulator.git
cd Internet-Protocol-Simulator
```

Baixe as dependências e atualize `go.mod` para que o projeto funcione corretamente:

```bash
go mod tidy
```

Abra duas janelas de terminal (ambas devem estar na raiz do respositório).
Inicie o servidor em uma delas:

```bash
go run main.go server
```

Inicie o cliente em outra:

```bash
go run main.go client
```

## Estrutura do Repositório

```bash
IP_sim/
├── client/                      # Código relacionado ao Cliente
│   ├── client.go                # Implementação principal do Cliente
│   └── gui/                     # Interface gráfica do Cliente
│       └── gui.go               # Implementação da GUI do Cliente
├── server/                      # Código relacionado ao Servidor
│   ├── server.go                # Implementação principal do Servidor
│   └── gui/                     # Interface gráfica do Servidor
│       └── gui.go               # Implementação da GUI do Servidor
├── link_layer/                  # Implementação da Camada de Enlace
│   ├── framing/                 # Protocolos de enquadramento
│   │   ├── char_count.go        # Contagem de caracteres
│   │   └── byte_insertion.go    # Inserção de bytes ou caracteres
│   ├── error_detection/         # Detecção de erros
│   │   ├── parity_bit.go        # Bit de paridade
│   │   └── crc.go               # CRC-32
│   └── error_correction/        # Correção de erros
│       └── hamming.go           # Código de Hamming
├── physical_layer/              # Implementação da Camada Física
│   ├── baseband_modulation/     # Modulações banda-base
│   │   ├── nrz_polar.go         # Modulação NRZ-Polar
│   │   ├── manchester.go        # Modulação Manchester
│   │   └── bipolar.go           # Modulação Bipolar
│   ├── carrier_modulation/      # Modulações por portadora
│   │   ├── ask.go               # Modulação ASK
│   │   ├── fsk.go               # Modulação FSK
│   │   └── qam_8.go             # Modulação 8-QAM
├── common/                      # Código compartilhado entre cliente e servidor
│   ├── utils/                   # Utilidades gerais
│   │   ├── encoding.go          # Funções auxiliares de codificação
│   │   └── math.go              # Funções matemáticas úteis
│   ├── communication/           # Sockets
│   │   ├── socket_client.go
│   │   ├── socket_server.go
│   │   └── protocol.go          # Protocolo de comunicação
│   └── constants.go             # Constantes compartilhadas
├── main.go                      # Ponto de entrada da aplicação
├── go.mod                       # Gerenciamento de dependências do Go
├── .gitignore
└── README.md
```

## Legenda para Issues
1.1.1  - Camada física - modulação digital
1.1.2  - Camada física - modulação por portadora
1.3    - Camada de enlace - protocolos de enquadramento de dados
1.4    - Camada de enlace - protocolos de detecção de erros
1.5    - Camada de enlace - protocolos de correção de erros
