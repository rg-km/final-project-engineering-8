import "../style/Footer.css"
import { FaFacebook, FaInstagram, FaTwitter } from "react-icons/fa";
import useStore from "../store/User";

export default function TeacherSection() {
    const { name } = useStore()
    return (
        <div className="footer-dark">
            <footer>
                <div className="container">
                    <div className="row">
                        <div className="col-sm-6 col-md-3 item">
                            <h3>Services</h3>
                            <ul>
                                <li><a href="/home">Home</a></li>
                                <li><a href="/tentang">Tentang</a></li>
                                <li><a href="/profile">Profile {name}</a></li>
                            </ul>
                        </div>
                        <div className="col-sm-6 col-md-3 item">
                            <h3>Tentang</h3>
                            <ul>
                                <li><a href="#">Company</a></li>
                                <li><a href="#">Team</a></li>
                                <li><a href="#">Careers</a></li>
                            </ul>
                        </div>
                        <div className="col-md-6 item text">
                            <h3>HALOGURU</h3>
                            <p>inovasi baru untuk mencari pendidik dengan sangat cepat.
                                Kegiatan belajar mengajar menjadi lebih mudah dengan menghadirkan tenaga didik ke rumah.</p>
                        </div>
                        <div className="col item social">
                            <a href="#"><i className="icon"><FaFacebook /></i></a>
                            <a href="#"><i className="icon"><FaTwitter /></i></a>
                            <a href="https://www.instagram.com/dindawr_/"><i className="icon"><FaInstagram /></i></a></div>
                    </div>
                    <p className="copyright">Kelompok 8 RGKM © 2022</p>
                </div>
            </footer>
        </div>

    )
}