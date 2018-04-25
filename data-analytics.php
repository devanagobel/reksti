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
        Redirect('DataAnalytics/dataAnalyticsStudent.php', false);
    }elseif(isset($_SESSION['userType']) == 'siswa'){
        //Nanti ini redirect ke halaman php khusus siswa
        Redirect('DataAnalytics/dataAnalyticsTeacher.php', false);
    }else{
        Redirect('index.html', false);
    }
?>
    