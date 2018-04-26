

$(function(){
    $('.button-confirm').click(function(){
        var data = {
            course_index: $("#course-index").val(),
            class_index: $("#class-index").val(),
            student_nim: $("#student-nim").val(),
        };

        $.ajax({
            url: '/API/attendance',
            type: 'POST',
            cache: false,
            data: JSON.stringify(data),
            dataType: "json",
            contentType: "application/json",
            success: function (response, textStatus, xhr) {
                if ((xhr.status != 200) || (response.course_index == undefined) || (response.class_index == undefined)
                    (response.student_nim == "undefined")){
                    alert("cannot put data")
                } else {
                    alert(response);
                }
            }
        })
    })
})