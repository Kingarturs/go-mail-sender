button = document.querySelector("#sendButton");

button.addEventListener("click", () => {
    to = document.querySelector("#to").value;
    subject = document.querySelector("#subject").value;
    content = document.querySelector("#content").value;

    
    if (to && subject && content) {
        url = "http://localhost:8000/send"
        $.ajax({
            url: url,
            method: 'POST',
            data: {
                to: to,
                subject: subject,
                content: content,
            },
            success: (res) => {
                Swal.fire({
                    position: 'bottom-end',
                    icon: 'success',
                    title: res.message,
                    showConfirmButton: false,
                    timer: 1000,
                    toast: true
                })
            },
            error: (error) => {
                console.log(error)
                Swal.fire({
                    position: 'bottom-end',
                    icon: 'error',
                    title: error.responseJSON.message,
                    showConfirmButton: false,
                    timer: 1000,
                    toast: true
                })
            }
        })
    }
})
