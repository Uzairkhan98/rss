# RSS Feed Management CLI Tool

This CLI tool allows users to manage RSS feeds, follow feeds, browse posts, and handle user authentication and registration.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
  - [login](#login)
  - [register](#register)
  - [reset](#reset)
  - [users](#users)
  - [agg](#agg)
  - [addfeed](#addfeed)
  - [feeds](#feeds)
  - [follow](#follow)
  - [following](#following)
  - [unfollow](#unfollow)
  - [browse](#browse)
- [Example Usage](#example-usage)

## Installation

This is a simple CLI tool that can be used to fetch and store the different RSS posts uploaded on a given website. In order to run this CLI tool it is expected that the user has Golang and PostgreSQL installed on their machines.

To install Golang kindly visit https://go.dev/dl/
To install PostgreSQL kindly visit https://www.postgresql.org/download/

After installing the above two softwares you also need to add a file called `.gatorconfig.json` which should contain your Postgres connection string in an object like the following:

`{
    db_url: {{your_connection_string}}
}`

After completing the above setup you can install this package using the command `go install github.com/uzairkhan98/rss`.

## Usage

Run the CLI tool using:

```bash
./rss-cli <command> [options]
```

## Commands

### login

Login as a specific user.

**Usage:**

```bash
./rss-cli login <username>
```

**Description:**  
Sets the current session to the specified user. The username should be already registered.

---

### register

Register a new user.

**Usage:**

```bash
./rss-cli register <username>
```

**Description:**  
Creates a new user with the specified username and sets it as the current session user.

---

### reset

Reset all users.

**Usage:**

```bash
./rss-cli reset
```

**Description:**  
Resets all registered users. This will delete all user data and exit the application.

---

### users

List all users.

**Usage:**

```bash
./rss-cli users
```

**Description:**  
Displays a list of all registered users. The currently logged-in user is marked as `(current)`.

---

### agg

Aggregate feed data at regular intervals.

**Usage:**

```bash
./rss-cli agg <duration>
```

**Description:**  
Starts collecting feeds at the specified interval. The `duration` must be in a format recognized by Go's `time.ParseDuration`, e.g., `1m` for one minute or `30s` for thirty seconds.

---

### addfeed

Add a new feed.

**Usage:**

```bash
./rss-cli addfeed <feed-name> <feed-url>
```

**Description:**  
Adds a new RSS feed with the specified name and URL, and automatically follows it under the current user.

---

### feeds

List all feeds.

**Usage:**

```bash
./rss-cli feeds
```

**Description:**  
Displays all available feeds.

---

### follow

Follow a specific feed.

**Usage:**

```bash
./rss-cli follow <feed-url>
```

**Description:**  
Follows the feed identified by the specified URL for the current user.

---

### following

List followed feeds.

**Usage:**

```bash
./rss-cli following
```

**Description:**  
Displays a list of all feeds that the current user is following.

---

### unfollow

Unfollow a specific feed.

**Usage:**

```bash
./rss-cli unfollow <feed-url>
```

**Description:**  
Unfollows the feed identified by the specified URL for the current user.

---

### browse

Browse posts from followed feeds.

**Usage:**

```bash
./rss-cli browse [limit]
```

**Description:**  
Displays posts from the feeds the current user is following. Optionally, specify a `limit` to control the number of posts retrieved (default is 2).

---

## Example Usage

1. **Register a user:**

   ```bash
   ./rss-cli register john_doe
   ```

2. **Login as the user:**

   ```bash
   ./rss-cli login john_doe
   ```

3. **Add a new feed:**

   ```bash
   ./rss-cli addfeed "Tech News" "https://technews.com/rss"
   ```

4. **Follow a feed:**

   ```bash
   ./rss-cli follow "https://technews.com/rss"
   ```

5. **Browse posts from followed feeds:**

   ```bash
   ./rss-cli browse 5
   ```

6. **Reset users (dangerous):**
   ```bash
   ./rss-cli reset
   ```
