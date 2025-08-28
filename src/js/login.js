document
  .getElementById("loginForm")
  .addEventListener("submit", async function (e) {
    e.preventDefault();

    const Email = document.getElementById("Email").value;
    const password = document.getElementById("password").value;

    if (!Email || !password) {
      showResult("Please fill in all fields", "red");
      return;
    }

    const data = { Email: Email, password: password };

    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });

      const result = await res.json();

      if (result.token) {
        showResult("Login successful", "green");
        localStorage.setItem("token", result.token);
        localStorage.setItem("role", result.role);
      } else {
        showResult(result.error || "Login failed", "red");
      }
    } catch (err) {
      showResult("Network error", "red");
    }
  });

function showResult(message, color) {
  const resultEl = document.getElementById("result");
  resultEl.innerText = message;
  resultEl.style.color = color;
}
