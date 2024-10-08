Here’s the English version of the general principles for naming methods effectively:

### 1. **Use Clear Verbs**

Methods typically perform actions, so their names should begin with a verb that clearly reflects what the method does.

- **Examples:**
  - `CreateUser()`: For creating a new user.
  - `UpdateOrderStatus()`: For updating the status of an order.
  - `DeleteBlogPost()`: For deleting a blog post.

### 2. **Method Names Should Be Specific and Descriptive**

The method name should be descriptive enough so that someone reading the code can easily understand what the method does without needing to look at the implementation.

- **Good Example:**
  - `FindUserByEmail()`: Finds a user based on their email.
  - `SendWelcomeEmail()`: Sends a welcome email.
- **Bad Example:**
  - `DoAction()`: Too generic and unclear.

### 3. **Maintain Consistent Naming Conventions**

Once you choose a certain word for an operation, be consistent with its usage throughout the codebase. For example, if you use `Find` for retrieval operations, use the same term across all similar methods.

- **Examples:**
  - `FindUserById()` → consistent use of "Find" for lookup methods.
  - `GetOrderDetails()` → consistent use of "Get" for retrieving details.

### 4. **Keep It Short but Informative**

Method names should be concise but still descriptive enough to convey their purpose. Avoid unnecessary words.

- **Good Example:**
  - `CalculateTotalPrice()`: Short and clear.
- **Bad Example:**
  - `DoTotalPriceCalculationForItemsInCart()`: Too long and verbose.

### 5. **Use Method Names that Reflect the Return Type**

If the method returns a certain value, the name should indicate what is being returned, not just the action performed.

- **Good Example:**
  - `GetUserEmail()`: Returns a user's email.
  - `FetchOrderList()`: Retrieves a list of orders.
- **Bad Example:**
  - `GetData()`: Too vague about what data is being retrieved.

### 6. **Avoid Using "And" in Method Names**

If a method name contains "and," it might indicate that the method is doing more than one thing. It's better to split it into multiple, more specific methods.

- **Bad Example:**
  - `SaveAndNotifyUser()`: This should likely be split into `SaveUser()` and `NotifyUser()`.

### 7. **Use Well-Known Naming Conventions**

Follow established naming conventions that are commonly used in the developer community. For example, in CRUD operations (Create, Read, Update, Delete), use standard words like:

- `Create`: For creating a new resource.
- `Get` or `Find`: For retrieving or reading data.
- `Update`: For updating existing data.
- `Delete`: For removing data.

- **Examples:**
  - `CreateBlogPost()`: For creating a new blog post.
  - `GetUserById()`: For retrieving a user by their ID.
  - `UpdateProductInfo()`: For updating product information.
  - `DeleteComment()`: For deleting a comment.

### 8. **Follow Asynchronous Naming Patterns**

If a method performs asynchronous operations, append a suffix that indicates the method is asynchronous, such as `Async` or `Promise`.

- **Examples:**
  - `SendEmailAsync()`: For asynchronously sending an email.
  - `FetchDataAsync()`: For asynchronously fetching data.

### 9. **Use Method Names that Reflect Business Processes**

If the method relates to a specific business process, its name should reflect that **business domain**.

- **Examples:**
  - `ProcessPayment()`: For processing a payment.
  - `CancelOrder()`: For canceling an order.
  - `ApproveTransaction()`: For approving a transaction.

### Examples of Good Method Names:

- `AddUser()`: Adds a user.
- `UpdateUserProfile()`: Updates a user's profile.
- `GetProductByID()`: Retrieves a product by its ID.
- `SendNotification()`: Sends a notification.
- `ProcessRefund()`: Processes a refund.

### Conclusion:

- **Use clear verbs** that describe the operation being performed.
- Ensure method names are **specific and descriptive**.
- Avoid overly **generic or ambiguous** naming.
- **Split methods** that perform multiple tasks into more specific methods.

By following these principles, method names will be clearer, more understandable, and accurately reflect their functionality within the business context.
