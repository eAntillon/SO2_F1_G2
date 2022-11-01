import Link from 'next/link';
import React from 'react';

const login = () => {
    return (
        <div className="flex items-center justify-center h-screen">
            <div className="flex flex-col">
                <h2 className="text-3xl font-semibold text-primary">Login</h2>
                <div className="mt-0 mb-1 divider" />
                <form className='max-w-xs'>
                    <label className="label">
                        <span className="label-text">Correo</span>
                    </label>
                    <input
                        type="email"
                        placeholder="hola@mail.com"
                        className="w-full max-w-xs input input-bordered"
                        required
                    />
                    <label className="label">
                        <span className="label-text">Contraseña</span>
                    </label>
                    <input
                        type="password"
                        placeholder="********"
                        className="w-full max-w-xs mb-4 input input-bordered"
                        required

                    />
                    <button className="w-full btn" type="submit">
                        Iniciar Sesion
                    </button>
                </form>
                <Link href="/registro" className="mt-2 text-sm link link-primary">Registrate</Link>
            </div>
        </div>
    );
};

export default login;
