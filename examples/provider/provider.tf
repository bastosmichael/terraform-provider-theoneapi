terraform {
  required_providers {
    theoneapi = {
      source = "bastosmichael/theoneapi"
      version = "0.1.0"
    }
  }
}

provider "theoneapi" {
  api_token = "YOUR_API_TOKEN_HERE"
}

resource "theoneapi_book" "example_book" {
  name = "The Fellowship of the Ring"
}

resource "theoneapi_movie" "example_movie" {
  name = "The Return of the King"
}

output "book_id" {
  value = theoneapi_book.example_book.id
}

output "movie_id" {
  value = theoneapi_movie.example_movie.id
}
