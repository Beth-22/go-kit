const express = require("express");
const { loginPage, dashboardPage } = require("../controllers/loginController");

const router = express.Router();

router.post("/login", loginPage);
router.get("/dashboard", dashboardPage);

module.exports = router;
