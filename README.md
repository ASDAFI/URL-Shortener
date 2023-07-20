# URL Shortener

This is a simple URL shortener application that utilizes MD5 hashing to create short links for URLs. The application is designed following Domain-Driven Design (DDD) principles and uses Redis as a cache to enhance performance.

## How it Works

The URL shortener provides two endpoints to handle shortening and retrieving original URLs:

1. **Shorten URL Endpoint**

   - Endpoint: `POST /short`
   - Request Body: `{"origin_link": "google.com"}`
   - Response: 
     ```json
     {
       "shortened_link": "1c383cd30b7c298ab50293adfecb7b18",
       "expires_at": "2023-07-20T03:55:33.257126Z",
       "message": "OK!"
     }
     ```

   When a `POST` request is made to the `/short` endpoint with the `origin_link`, the application generates an MD5 hash for the URL, creating a shortened link. The response includes the `shortened_link` and an `expires_at` timestamp indicating the expiration date of the short link. Additionally, a "OK!" message is returned to indicate the successful creation of the short link.

2. **Retrieve Original URL Endpoint**

   - Endpoint: `POST /origin`
   - Request Body: `{"shortened_link": "1c383cd30b7c298ab50293adfecb7b18"}`
   - Response: 
     ```json
     {
       "origin_link": "google.com",
       "expires_at": "2023-07-20T03:55:33Z",
       "message": "OK!"
     }
     ```

   When a `POST` request is made to the `/origin` endpoint with the `shortened_link`, the application looks up the original URL associated with the given shortened link and returns it. The response includes the `origin_link` and an `expires_at` timestamp indicating the expiration date of the original link. A "OK!" message is also included in the response.

## Requirements

To run the URL shortener application, you need the following:

- Go programming language (version X.X.X)
- Redis server (for caching)

## Installation and Setup

1. Clone the repository
2. Setup Deployments

Before running the URL shortener, ensure that you have set up the necessary deployments, such as setting up the Redis server or any other configurations required for your specific environment using this command:
```
make developements
```


3. Build Go Files

Use the following command to build the Go files for the URL shortener:
```
make url-shortener
```

4. Migrate Database
Run the migration to set up the database schema for the URL shortener:
```
./url-shortener migrate
```

5. Run Server

Start the URL shortener server:
```
./url-shortener serve
```

## Note

Feel free to contribute and improve the application. Happy URL shortening!




