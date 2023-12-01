**Product Catalogue RESTful API**
=

**Overview** 
=
- This project aims to create a RESTful API for a product catalogue with a focus on managing product information and reviews. The API allows clients to perform various operations related to products and reviews.
---
**How To Run**
=
- Clone project
- Run Command ``docker-compose up``
- *this repo still have a configuration deployment problem. Will fix ASAP
- During resolve deployment issue. After running ``docker-compose up``. It possible to using ``localhost:8080``as local host & port
---
**Scope**
=
- **Product** 
  - Attributes: SKU, title, description, category, images, weight, price.
  - Uniqueness: SKU must be unique for each product.
  - Product Image Description: Short description available for each product image.
- **Product Reviews**
  - Attributes: rating, review comment. review date
---
**Functional Requirements**
=
**Product Operations**
- Create Product
- Update Product: Ability to update SKU, title, description, and category.
- Get Product by ID
- Search Products 
  - Criteria: SKU, title, category.
  - Sort products by date (newest or oldest).
  - Sort products based on rating (highest or lowest).
  - Pagination support

**Product Reviews Operations**
- Add Product Reviews and Ratings (Products can have multiple reviews and ratings)
---
**API Endpoints**
=
**Product Endpoints**
- Create Product 
  - POST /api/v1/product
``` 
  REQUEST BODY
  {
    "sku": "123456790",
    "title": "Aqua 250ml",
    "description": "Air Mineral Berkualitas",
    "category": "Air Mineral",
    "imgURL": "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
    "weight": 0.70,
    "price": 11000
  }
```
- Update Product 
  - PATCH /api/v1/product/:sku
``` 
  PATH: api/v1/product/123456791
  REQUEST BODY
  {
    "sku": "123456790",
    "title": "Aqua 250ml",
    "description": "Air Mineral Normal",
    "category": "Air Mineral"
  }
```
- Get Product by SKU 
  - GET /api/v1/product/:sku ```PATH: api/v1/product/123456790```
- Search Products With All Filters
  - GET /api/v1/product 
    - Parameters: 
      - sku (non-mandatory)
      - title (non-mandatory)
      - category (non-mandatory)
      - sort (non-mandatory)
      - page_index (non-mandatory)
      - page_size (non-mandatory)
      - sort (non-mandatory)
        - rating (sort=1:asc or sort=2:desc)
        - product_date (sort=2:asc or sort=2:desc)
  - for example: ``/api/v1/product?category=Air Mineral&sort=2:asc``

**Product Reviews Endpoints**
- Add Product Reviews and Ratings 
  - POST /api/v1/productreview
``` 
  REQUEST BODY
  {
    "productSKU": "123456790",
    "rating": 1,
    "reviewComment": "Buruk"
  }
```