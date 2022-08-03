import axios from "axios";


export async function DownloadFile(file) {
        const response = await axios({
            url: "http://localhost:8080/download?name=" + file,
            method: 'GET',
            responseType: 'blob', // important
        })
        if (response.status === 200) {
            const url = window.URL.createObjectURL(new Blob([response.data]));
            const link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', file); //or any other extension
            document.body.appendChild(link);
            link.click();
            link.remove()
        }
}