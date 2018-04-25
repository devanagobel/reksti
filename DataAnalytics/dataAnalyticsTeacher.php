<?php
    function Redirect($url, $permanent = false)
    {
        header('Location: ' . $url, true, $permanent ? 301 : 302);

        exit();
    }
    //ini jadi middleware
    session_start();
    if(isset($_SESSION['userType']) == 'dosen') {
        //Nanti ini ganti redirect ke halaman php khusus dosen
        Redirect('../DataAnalytics/dataAnalyticsStudent.php', false);
    }elseif(isset($_SESSION['userType']) == 'siswa'){
        //Nanti ini redirect ke halaman php khusus siswa
        Redirect('../DataAnalytics/dataAnalyticsTeacher.php', false);
    }else{
        Redirect('../index.html', false);
    }
?>
    
<!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Document</title>
    </head>
    <body>
        <div class="data-analytic">
            <!-- ChangeClass with php code -->
            <?php
                $skill = "demo";
                echo '<div class= '. $skill . ' data-percent="65"></div>';
            ?>
        </div>

        <script src="http://ajax.googleapis.com/ajax/libs/jquery/2.1.3/jquery.min.js"></script>
        <script src="../js/jquery.circlechart.js"></script>
        <script>

            $('.demo').percentcircle({
            animate : true,
            diameter : 100,
            guage: 3,
            coverBg: '#fff',
            bgColor: '#efefef',
            fillColor: '#46CFB0',
            percentSize: '18px',
            percentWeight: 'normal'
            });		
        </script>
    </body>



    </html>
