import { useCallback, useEffect } from 'react';
import { useGetState } from './useGetState';
// import { store } from '../utils/store';
// import { debounce } from '../utils';

export const useConfig = (key, defaultValue, options = {}) => {
    const [property, setPropertyState, getProperty] = useGetState(null);
    // const { sync = true } = options;

    // 同步到Store (State -> Store)
    // const syncToStore = useCallback(
    //     debounce((v) => {
    //         store.set(key, v);
    //         store.save();
    //         // emit(`${key}_changed`, v);
    //     }),
    //     []
    // );

    // 同步到State (Store -> State)
    const syncToState = useCallback((v) => {
        if (v !== null) {
            setPropertyState(v);
        } else {
            if (getProperty() === null) {
                setPropertyState(defaultValue);
            } else {
                setPropertyState(v);
            }
            // store.get(key).then((v) => {
            //     if (v === null) {
            //         setPropertyState(defaultValue);
            //         store.set(key, defaultValue);
            //         store.save();
            //     } else {
            //         setPropertyState(v);
            //     }
            // });
        }
    }, []);

    const setProperty = useCallback((v, forceSync = false) => {
        setPropertyState(v);
        // const isSync = forceSync || sync;
        // isSync && syncToStore(v);
    }, []);

    // 初始化
    useEffect(() => {
        syncToState(null);
        if (key.includes('[')) return;
        // const unlisten = listen(`${key}_changed`, (e) => {
        //     syncToState(e.payload);
        // });
        // return () => {
        //     unlisten.then((f) => {
        //         f();
        //     });
        // };
    }, []);

    return [property, setProperty, getProperty];
};
