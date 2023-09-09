package theoneapi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBook() *schema.Resource {
	return &schema.Resource{
		Create: resourceBookCreate,
		Read:   resourceBookRead,
		Update: resourceBookUpdate,
		Delete: resourceBookDelete,

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"author": {
				Type:     schema.TypeString,
				Required: true,
			},
			"published_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Continue with other fields as per your OpenAPI spec
		},
	}
}

func resourceBookCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	book := &Book{
		Title:  d.Get("title").(string),
		Author: d.Get("author").(string),
		// and so on...
	}

	createdBook, err := client.CreateBook(book)
	if err != nil {
		return err
	}

	d.SetId(createdBook.ID)
	return resourceBookRead(d, m)
}

func resourceBookRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	book, err := client.GetBook(d.Id())
	if err != nil {
		return err
	}

	d.Set("title", book.Title)
	d.Set("author", book.Author)
	// Continue with other fields

	return nil
}

func resourceBookUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	book := &Book{
		Title:  d.Get("title").(string),
		Author: d.Get("author").(string),
		// and so on...
	}

	_, err := client.UpdateBook(d.Id(), book)
	if err != nil {
		return err
	}

	return resourceBookRead(d, m)
}

func resourceBookDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteBook(d.Id())
	if err != nil {
		return err
	}

	return nil
}
