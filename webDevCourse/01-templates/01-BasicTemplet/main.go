package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// HTML code is from https://www.geeksforgeeks.org/html-layout/
	data := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Page Layout</title>
	<style>
	.head1 {
		font-size:40px;
	color:#009900;
		font-weight:bold;
	}
	.head2 {
		font-size:17px;
		margin-left:10px;
		margin-bottom:15px;
	}
	body {
		margin: 0 auto;
		background-position:center;
		background-size: contain;
	}

	.menu {
		position: sticky;
		top: 0;
		background-color: #009900;
		padding:10px 0px 10px 0px;
		color:white;
		margin: 0 auto;
		overflow: hidden;
	}
	.menu a {
		float: left;
		color: white;
		text-align: center;
		padding: 14px 16px;
		text-decoration: none;
		font-size: 20px;
	}
	.menu-log {
		right: auto;
		float: right;
	}
	footer {
		width: 100%;
		bottom: 0px;
		background-color: #000;
		color: #fff;
		position: absolute;
		padding-top:20px;
		padding-bottom:50px;
		text-align:center;
		font-size:30px;
		font-weight:bold;
	}
	.body_sec {
		margin-left:20px;
	}
	</style>
	</head>

	<body>

	<!-- Header Section -->
	<header>
	<div class="head1">GeeksforGeeks</div>
	<div class="head2">A computer science portal for geeks</div>
	</header>

	<!-- Menu Navigation Bar -->
	<div class="menu">
	<a href="#home">HOME</a>
	<a href="#news">NEWS</a>
	<a href="#notification">NOTIFICATIONS</a>
	<div class="menu-log">
	<a href="#login">LOGIN</a>
	</div>
	</div>

	<!-- Body section -->
	<div class = "body_sec">
	<section id="Content">
	<h3>Content section</h3>
	</section>
	</div>

	<!-- Footer Section -->
	<footer>Footer Section</footer>
	</body>
	</html>
	`

	nf, err := os.Create("index.gohtml")
	if err != nil {
		log.Fatal("error making file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(data))
}
