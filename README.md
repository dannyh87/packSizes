# GymShark Test - Pack Sizes
### Brief

Imagine for a moment that one of our product lines ships in various pack sizes:
- 250 Items
- 500 Items
- 1000 Items
- 2000 Items
- 5000 Items


Our customers can order any number of these items through our website, but they
will always only be given complete packs.

1. Only whole packs can be sent. Packs cannot be broken open.

2. Within the constraints of Rule 1 above, send out no more items than
necessary to fulfil the order.
3. Within the constraints of Rules 1 &amp; 2 above, send out as few packs as
possible to fulfil each order.

### Algorithm
example : 
pack request : 20001
pack sizes : 500, 250, 1000, 2500, 5000

Basic outline :

Sort the pack sizes to ensure they are in size order (user may not input them in this way)
Reduce the pack request by looping through the pack sizes in reverse, removing the largest pack possible each time.

Using the example above

4 x 5000
1 x 250

Add each pack selected to a results array.

### How to run
```bash
go build
./dh_pack_size
```
Port 8080 is used for the frontend:
http://localhost:8080/
### Approach
As I'm not familiar with go this was originally written in node.js firstly solving the algorthim, and trying to resolve any inital bugs to match the rules outlined in the challenge.
The Node.js code was then translated to Go using material which was available online.

Postman was utilised to ensure the API was working as expected
Front end was added using HTML, CSS and Vanilla JS.
### Improvements
- Improve the styling to match Gymshark branding
- Increase data which is provided to the end user such as total pack numbers and wastage
- Keep a track of previous requests and responses
- Add testing