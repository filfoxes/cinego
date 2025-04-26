# Cinema Room Manager

A command-line application for managing a cinema room's seating and ticket sales.

## Features

- **Seating Visualization**: Display the current state of cinema seats
- **Ticket Purchase**: Book seats with dynamic pricing based on room size and seat location
- **Statistics**: View detailed statistics including:
  - Number of purchased tickets
  - Percentage of occupancy
  - Current income
  - Total potential income

## Pricing Rules

- For rooms with less than 60 seats:
  - All tickets are $10
- For rooms with 60 or more seats:
  - Front half rows: $10 per ticket
  - Back half rows: $8 per ticket

## Usage

1. Start the program and enter:
   - Number of rows in the cinema
   - Number of seats in each row

2. Use the menu to:
   - Show the seating arrangement (Option 1)
   - Buy a ticket (Option 2)
   - View statistics (Option 3)
   - Exit the program (Option 0)

## Seat Legend
- `S`: Available seat
- `B`: Booked seat

## Input Validation

The program includes validation for:
- Row and seat numbers (must be within valid range)
- Already purchased tickets
- Menu option selection

## Example Output
