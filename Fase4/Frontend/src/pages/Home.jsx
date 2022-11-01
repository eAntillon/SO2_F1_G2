import { useEffect, useState } from 'react';
import Memsim from '../components/Memsim';
import { useNavigate } from "react-router-dom";

export default function Home() {
    const router = useNavigate();
    const [auth, setAuth] = useState(false);
    useEffect(() => {
        const auth = localStorage.getItem('auth');
        if (auth) {
            setAuth(true);
        }
    }, []);

    const logout = () => {
        localStorage.removeItem('auth');
        router('/login');
    };

    return (
        <div>
            <main className="flex items-center justify-center h-screen">
                <div className="min-h-screen hero bg-base-200">
                    <div className="text-center hero-content">
                        <div className="max-w-lg">
                            <h1 className="text-4xl font-bold">
                                MetaOS - Fase 4
                            </h1>
                            {auth ? (
                                <Memsim logout={logout} />
                            ) : (
                                <div className="mt-5">
                                    <p>No has iniciado sesion :(</p>
                                    <button
                                        className="mt-2 link link-primary"
                                        onClick={() => router('/login')}
                                    >
                                        Iniciar sesion
                                    </button>
                                </div>
                            )}
                        </div>
                    </div>
                </div>
            </main>
        </div>
    );
}
