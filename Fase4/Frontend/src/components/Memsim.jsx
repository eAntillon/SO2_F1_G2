import axios from 'axios';
import React, { useState } from 'react';
import Code from './Code';

const Memsim = ({ logout }) => {
    const [valores, setValores] = useState({
        ciclos: '',
        unidades: 'u1,u2,u3',
    });
    const [memsimData, setMemsimData] = useState([]);

    const handleChange = (e) => {
        setValores({
            ...valores,
            [e.target.name]: e.target.value,
        });
    };

    const ejecutar = () => {
        if (valores.unidades === '' || valores.ciclos === '') {
            console.log('Error');
            return;
        }
        setMemsimData([]);
        console.log('Ejecutando', valores);
        axios
            .post('api/memsim', {
                ciclos: parseInt(valores.ciclos),
                unidades: valores.unidades,
            })
            .then(function (response) {
                console.log(response);
                if (response.status === 200) {
                    console.log('Ejecucion exitosa');
                    setMemsimData(response.data);
                }
            })
            .catch(function (error) {
                console.log(error);
            });
    };

    return (
        <div>
            <h1 className="mt-5 mb-2 text-2xl font-bold text-primary">
                Memsim
            </h1>
            <div className="flex items-end gap-4">
                <div>
                    <label className="label">
                        <span className="label-text">Ciclos</span>
                    </label>
                    <input
                        type="number"
                        placeholder="No. Ciclos"
                        className="w-full input input-bordered input-sm"
                        name="ciclos"
                        value={valores.ciclos}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label className="label">
                        <span className="label-text">Unidades</span>
                    </label>
                    <input
                        type="text"
                        placeholder="u1,u2,u3"
                        className="w-full input input-bordered input-sm"
                        name="unidades"
                        value={valores.unidades}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <button
                        className="btn btn-primary btn-sm"
                        onClick={ejecutar}
                    >
                        Simular
                    </button>
                </div>
            </div>
            <div className="m-0 mb-2 divider"></div>

            <Code data={memsimData} />

            <button onClick={logout} className="btn btn-error btn-sm">
                Cerrar sesión
            </button>
        </div>
    );
};

export default Memsim;
