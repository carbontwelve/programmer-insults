<?php

$insultsFile = __DIR__ . DIRECTORY_SEPARATOR . "insults.txt";

/**
 * Function for grabbing a random line from a file
 *
 * @author Mohamed Nuur
 * @link http://stackoverflow.com/a/12119028/1225977
 * @param string $fileName
 * @param int $maxLineLength
 * @return string
 */
function rand_line($fileName, $maxLineLength = 4096) {
    $handle = @fopen($fileName, "r");
    if ($handle) {
        $random_line = null;
        $line = null;
        $count = 0;
        while (($line = fgets($handle, $maxLineLength)) !== false) {
            $count++;
            // P(1/$count) probability of picking current line as random line
            if(rand() % $count == 0) {
              $random_line = $line;
            }
        }
        if (!feof($handle)) {
            echo "Error: unexpected fgets() fail\n";
            fclose($handle);
            return null;
        } else {
            fclose($handle);
        }
        return $random_line;
    }
}

?><!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">

  <title>Programmer Insults!</title>
  <meta name="description" content="Insults for programmers">
  <meta name="author" content="Simon Dann">

  <link rel="stylesheet" href="vendor/HTML5-Reset/assets/css/reset.css">
  <link href="//fonts.googleapis.com/css?family=Fira+Mono:400,700" rel="stylesheet" type="text/css">
  <link rel="stylesheet" href="styles.css?v=1.0">
</head>

<body>
    <main>
        <h1><?php echo rand_line($insultsFile); ?></h1>
        <br><br>
        <small><a href="https://www.programmerinsults.com/">More?</a></small>
    </main>
    <footer>Insults sourced from various places. <a href="https://github.com/carbontwelve/programmer-insults">src</a></footer>
</body>
</html>
