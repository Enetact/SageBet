
# 🧠 SageBet

SageBet is a sports betting analytics platform that leverages statistical models to predict outcomes and optimize betting strategies. Built with Go, it integrates various components to collect data, compute probabilities, and apply betting algorithms.

## 📋 Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## 🚀 Features

- **Data Collection**: Aggregates sports data for analysis.
- **Probability Modeling**: Utilizes Poisson distributions and other statistical methods to predict outcomes.
- **Betting Strategy**: Implements the Kelly criterion to optimize bet sizing.
- **Modular Design**: Structured components for bettors, models, and data collectors.

## 🛠️ Getting Started

### Prerequisites

- Go 1.16 or higher
- Git

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Enetact/SageBet.git
   cd SageBet
   ```

2. **Build the application**:

   ```bash
   go build -o sagebet
   ```

## 💡 Usage

1. **Run the application**:

   ```bash
   ./sagebet
   ```

2. **Configure settings**:

   Modify configuration parameters within the source code or implement a configuration file as needed.

## 📁 Project Structure

```
SageBet/
├── ai_bettor.go          # Core betting logic
├── collector.go          # Data collection utilities
├── kelly.go              # Kelly criterion implementation
├── main.go               # Entry point of the application
├── models.go             # Data models and structures
├── poisson.go            # Poisson distribution calculations
├── vector_client.go      # Client for vector operations
├── go.mod                # Go module file
├── go.sum                # Go module checksums
└── README.md             # Project documentation
```

## 🤝 Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/YourFeature`
3. Commit your changes: `git commit -m 'Add YourFeature'`
4. Push to the branch: `git push origin feature/YourFeature`
5. Open a pull request.

For major changes, please open an issue first to discuss what you would like to change.

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 📬 Contact

For questions or suggestions, please open an issue on the [GitHub repository](https://github.com/Enetact/SageBet/issues).
