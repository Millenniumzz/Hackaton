document
  .getElementById("registerForm")
  .addEventListener("submit", async function (e) {
    e.preventDefault();

    const password = document.getElementById("password").value;
    const confirmPassword = document.getElementById("confirmPassword").value;

    if (password.length < 8) {
      document.getElementById("result").innerText =
        "Password must be at least 8 characters!";
      document.getElementById("result").style.color = "red";
      return;
    }

    if (password !== confirmPassword) {
      document.getElementById("result").innerText = "Passwords do not match!";
      document.getElementById("result").style.color = "red";
      return;
    }

    const data = {
      username: document.getElementById("username").value,
      email: document.getElementById("email").value,
      password: password,
    };

    try {
      const res = await fetch("http://localhost:8080/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });

      const result = await res.json();
      document.getElementById("result").innerText =
        result.message || result.error;
      document.getElementById("result").style.color = result.message
        ? "green"
        : "red";
    } catch (err) {
      document.getElementById("result").innerText = "Network error";
      document.getElementById("result").style.color = "red";
    }
  });
