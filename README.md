# Groupie Tracker

* Groupie Tracker is a web application that allows users to explore information about various bands and artists, including their history, members, concert locations, and event dates. 

* This project aims to create a user-friendly platform that visualizes data from a given API, providing an interactive and engaging experience.

## Table of Contents

- [Objectives](#objectives)
- [Project Overview](#project-overview)
- [API Structure](#api-structure)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Best Practices](#best-practices)
- [Contributing](#contributing)
- [License](#license)

## Objectives

The main objective of the Groupie Tracker project is to:

- Create a website that displays information about bands and artists using data from a provided API.
- Develop an interactive feature that triggers an action on the client side, making a request to the server to receive a response, enhancing user engagement with real-time data.

## Project Overview

The Groupie Tracker application consists of:

1. A backend written in Go that handles all server-side logic.
2. A frontend interface that presents data to the user in a user-friendly manner.
3. Data visualizations to showcase information about artists, concert locations, and event dates.
4. Interactive client-server communication to dynamically fetch and display data based on user actions.

## API Structure

The provided API consists of four main parts:

1. **Artists**: Contains information about bands and artists, including their name, image, start year, first album date, and members.
2. **Locations**: Lists the last and/or upcoming concert locations of the artists.
3. **Dates**: Provides details of the last and/or upcoming concert dates.
4. **Relation**: Links artists, dates, and locations, establishing connections between them.

## Features

- Display artist profiles with details such as name, image, members, and history.
- Show upcoming and past concert locations and dates.
- Interactive data visualizations for a better user experience.
- Client-server communication to fetch data dynamically based on user input.
- Responsive design for optimal viewing on various devices.

## Technologies Used

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS
- **API**: Provided by the project with endpoints for artists, locations, dates, and relations.

## Installation

1. Clone the repository:

   ```bash
   git clone https://learn.zone01kisumu.ke/git/fonditi/groupie-tracker.git
   cd groupie-tracker
   ```

2. Install Go dependencies:

```bash
go mod tidy

```

3. Run the Server

```bash
go run .
```

4. Open your browser and navigate to `http://localhost:8080.`

## Usage

* Browse the homepage to see a list of artists and bands.
* Click on an artist to view detailed information about their history, members, and concerts.
* Use the search feature to find specific artists or concerts.
* Interact with the visual elements to explore concert locations and dates.
  
## Shortcut Guide

- **`Ctrl + F`** — Focus Search
- **`Alt + H`** — Home
- **`Alt + ArrowLeft`** or **`Alt + B`** — Back
- **`Alt + Shift + ?`** — Help


## API Endpoints

The application interacts with the following API endpoints:

* **/artists:** Fetches details about artists and bands.
* **/locations:** Retrieves concert locations.
* **/dates:** Gets concert dates.
* **/relations:** Establishes relations between artists, locations, and dates.


## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (git checkout -b feature-branch).
3. Make your changes.
4. Commit your changes (git commit -m 'Add new feature').
5. Push to the branch (git push origin feature-branch).
6. Open a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributors

[Benard Opiyo](https://github.com/benardopiyo) 
[Fena] (https://github.com/konditi1)
