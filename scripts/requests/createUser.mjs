const params = {
  email: "user-test@example.com",
  password: "foobar2000"
};

const res = await fetch("http://localhost:3000/users", {
  method: "POST",
  body: JSON.stringify(params),
  headers: {
    "content-type": "application/json"
  }
});

const json = await res.text();

console.log(json);
