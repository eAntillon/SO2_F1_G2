import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const Registro = () => {
    const router = useNavigate();

    const [loading, setLoading] = useState(false);

    const [form, setForm] = useState({
        email: 'hola@gmail.com',
        password: 'asdfadf',
        confirmPassword: 'asdfadf1',
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
        if (form.password !== form.confirmPassword) {
            console.log('Las contraseñas no coinciden');
            return;
        }
        axios
            .post('api/registro', {
                correo: form.email,
                password: form.password,
            })
            .then(function (response) {
                console.log(response);
                if (response.status === 200) {
                    console.log('Registro exitoso');
                    router('/login');
                }
            })
            .catch(function (error) {
                console.log(error);
            });
    };

    return (
        <div className="flex items-center justify-center h-screen">
            <div className="flex flex-col">
                <h2 className="text-3xl font-semibold text-primary">
                    Registro
                </h2>
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
                        className="w-full max-w-xs input input-bordered"
                        value={form.password}
                        name="password"
                        onChange={handleChange}
                        required
                    />
                    <label className="label">
                        <span className="label-text">Confirmar Contraseña</span>
                    </label>
                    <input
                        type="password"
                        placeholder="********"
                        className="w-full max-w-xs mb-4 input input-bordered"
                        value={form.confirmPassword}
                        name="confirmPassword"
                        onChange={handleChange}
                        required
                    />
                    <button
                        className={
                            'w-full btn' + (loading ? ' loading disabled' : '')
                        }
                        type="submit"
                    >
                        Registrarse
                    </button>
                </form>
                <a href="/login" className="mt-2 text-sm link link-primary">
                    Iniciar Sesion
                </a>
            </div>
        </div>
    );
};

export default Registro;
