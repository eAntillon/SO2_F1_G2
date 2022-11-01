import axios from 'axios';
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const router = useNavigate();
    const [loading, setLoading] = useState(false);
    const [form, setForm] = useState({
        email: 'hola@gmail.com',
        password: 'pass',
    });

    const handleChange = (e) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value,
        });
    };
    const handleSubmit = (e) => {
        e.preventDefault();
        setLoading(true);
        axios
            .post('api/login', {
                correo: form.email,
                password: form.password,
            })
            .then(function (response) {
                console.log(response);
                if (response.status === 200) {
                    console.log(response);
                    console.log('Registro exitoso');
                    if (
                        response.data.status &&
                        response.data.status === 'authorized'
                    ) {
                        localStorage.setItem('auth', true);
                        router('/');
                    }
                }
            })
            .catch(function (error) {
                console.log(error);
            });
    };

    return (
        <div className="flex items-center justify-center h-screen">
            <div className="flex flex-col">
                <h2 className="text-3xl font-semibold text-primary">Login</h2>
                <div className="mt-0 mb-1 divider" />
                <form className="max-w-xs" onSubmit={handleSubmit}>
                    <label className="label">
                        <span className="label-text">Correo</span>
                    </label>
                    <input
                        type="email"
                        placeholder="hola@mail.com"
                        className="w-full max-w-xs input input-bordered"
                        value={form.email}
                        name="email"
                        onChange={handleChange}
                        required
                    />
                    <label className="label">
                        <span className="label-text">Contraseña</span>
                    </label>
                    <input
                        type="password"
                        placeholder="********"
                        className="w-full max-w-xs mb-4 input input-bordered"
                        value={form.password}
                        name="password"
                        onChange={handleChange}
                        required
                    />
                    <button
                        className={
                            'w-full btn' + (loading ? ' loading disabled' : '')
                        }
                        type="submit"
                    >
                        Iniciar Sesion
                    </button>
                </form>
                <a href="/registro" className="mt-2 text-sm link link-primary">
                    Registrate
                </a>
            </div>
        </div>
    );
};

export default Login;
