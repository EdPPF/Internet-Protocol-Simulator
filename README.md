```
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
│   ├── protocol.go              # Protocolo comum de comunicação
│   └── constants.go             # Constantes compartilhadas
├── main.go                      # Ponto de entrada da aplicação
├── go.mod                       # Gerenciamento de dependências do Go
├── .gitignore
└── README.md
```
