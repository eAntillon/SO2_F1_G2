import React from 'react';

const Code = ({ data }) => {
    return (
        <div className="mb-4 text-left mockup-code">
            <div className='h-64 overflow-visible overflow-y-auto'>
                {Object.values(data)
                    .join(',')
                    .split(',')
                    .map((hilo, index) => {
                        return (
                            <pre data-prefix={index + 1} key={index}>
                                <code className='text-xs'>{hilo}</code>
                            </pre>
                        );
                    })}
            </div>
        </div>
    );
};

export default Code;
