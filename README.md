# Telegram Mental Health Bot

## Description

This is a Telegram bot designed to provide mental health support by connecting users with available experts. The bot facilitates private sessions between users and mental health professionals, includes crisis detection features, and ensures secure message routing.

## Features

- User session management with mental health experts
- Automatic crisis word detection for immediate intervention
- Secure message routing between users and experts
- Bot filtering to prevent automated interactions
- Session lifecycle management (start, end, availability tracking)
- Command-based interface with /start, /end, and /help commands

## Prerequisites

- Go 1.22.3 or later
- A Telegram Bot Token (obtained from @BotFather on Telegram)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/aman1521coder/telegrammentalhealth.git
   cd telegrammentalhealth
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Build the application:
   ```
   go build
   ```

## Configuration

1. Obtain a bot token from Telegram's BotFather.
2. Set the environment variable for the bot token:
   ```
   export TELEGRAM_BOT_TOKEN="your_bot_token_here"
   ```

## Usage

1. Run the bot:
   ```
   ./telegrammentalhealth
   ```

2. Available commands:
   - `/start`: Initiate a session with an available expert
   - `/end`: End the current session
   - `/help`: Display available commands

The bot will automatically:
- Detect crisis-related keywords in messages
- Route messages between users and assigned experts
- Manage expert availability and session states

## Project Structure

- `main.go`: Entry point of the application with goroutines for updates and message handling
- `telegram.go`: Telegram Bot API integration, message handling, and command processing
- `client.go`: Session and expert management logic
- `utils.go`: Utility functions including crisis detection
- `go.mod`: Go module dependencies

## Contributing

Contributions are welcome. Please follow these steps:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Disclaimer

This bot is not a substitute for professional medical advice, diagnosis, or treatment. Always seek the advice of qualified health providers with questions about a medical condition.