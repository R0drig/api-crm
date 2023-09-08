# API Documentation

Welcome to the API documentation for your application. This document outlines the available functionalities and endpoints provided by the API.

## Authentication

To access protected endpoints, you need to include an `Authorization` header with a valid API key.

### Example Authorization Header

## Endpoints

### Register a User

Register a new user.

- **URL:** `/register`
- **Method:** `POST`
- **Request Body:**
  - `name` (string): User's name.
  - `email` (string): User's email address.
  - `passwd` (string): User's password.

### User Login

Authenticate and log in a user.

- **URL:** `/login`
- **Method:** `POST`
- **Request Body:**
  - `email` (string): User's email address.
  - `passwd` (string): User's password.

### User Profile

Retrieve the user's profile information.

- **URL:** `/auth/user`
- **Method:** `GET`

### Create a Lead

Create a new lead associated with the authenticated user.

- **URL:** `/auth/register-lead`
- **Method:** `POST`
- **Request Body:**
  - `name` (string): Lead's name.
  - `email` (string): Lead's email address.
  - `phone` (string): Lead's phone number.
  - `notes` (string): Additional notes for the lead.

### List All User Leads

Retrieve a list of all leads associated with the authenticated user.

- **URL:** `/auth/leads`
- **Method:** `GET`

### Update a Lead

Update an existing lead.

- **URL:** `/auth/update-lead/:id`
- **Method:** `PATCH`
- **Request Body:** Include fields you want to update.
  - `name` (string): New lead name.
  - `email` (string): New lead email.
  - `phone` (string): New lead phone number.
  - `notes` (string): New lead notes.

### Delete a Lead

Delete an existing lead.

- **URL:** `/auth/update-lead/:id`
- **Method:** `DELETE`

## Response Format

All responses are in JSON format and include appropriate status codes. In case of errors, the response will include an error message.

### Example Response (Success)
```json
{
  "message": "Operation successful",
  "data": { ... }
}
