These APIs provide the core functionality for our Weather Service, allowing users to register, log in, query current weather, access historical data, manage their historical weather records, and view/update their profile information. The use of JWT tokens for authentication and authorization is essential for securing these APIs.

List of APIs for the Weather Service:

## Weather Service APIs

### User Registration API
- **Description**: Allows users to register with the application by providing a username (email), password, and date of birth.
- **Flow**:
  - The API takes user registration data.
  - It validates the data and creates a new user record in the database.
  - It returns a success message or an error if registration fails.

### User Login API
- **Description**: Allows registered users to log in with their username (email) and password.
- **Flow**:
  - The API takes user login credentials.
  - It validates the credentials and generates a JWT token if they are correct.
  - It returns the JWT token & user ID or an authentication error.

### Get Current Weather API
- **Description**: Allows authenticated users to query the current weather conditions for a specific city.
- **Flow**:
  - The user sends an authenticated request with a city name.
  - The server queries an external weather API (e.g., OpenWeatherMap) for the current weather data.
  - It stores the query and weather data in the database for historical tracking.
  - The API returns the current weather data.

### Get Weather History API
- **Description**: Allows authenticated users to retrieve their historical weather queries.
- **Flow**:
  - The user sends an authenticated request with parameters like index(page no) and recordLength (list size).
  - The server queries the database to retrieve the user's historical weather queries.
  - It returns a list of weather query records within the specified date range.

### Delete Weather History Record API
- **Description**: Allows authenticated users to delete a specific historical weather record by its ID.
- **Flow**:
  - The user sends an authenticated request with the ID of the record to be deleted.
  - The server verifies the user's identity and the record's ownership.
  - It deletes the record from the database and returns a success message or an error.

### Bulk Delete Weather History Records API
- **Description**: Allows authenticated users to delete multiple historical weather records by providing a list of record IDs.
- **Flow**:
  - The user sends an authenticated request with a list of record IDs to be deleted in bulk.
  - The server verifies the user's identity and the records' ownership.
  - It deletes the records from the database and returns a success message or an error.

### User Profile API (BACKLOG : FUTURE SCOPE)
- **Description**: Allows users to view and update their profile information, including username (email) and date of birth.
- **Flow**:
  - The user sends an authenticated request to view or update their profile.
  - The server retrieves or updates the user's profile data in the database.
  - It returns the updated profile information or an error message.
