import { useEffect, useState } from 'react';
import axios from 'axios';

interface Block {
  BlockHeight: string;
  Receipient: string;
  Reward: string;
  Size: string;
  GasUsed: string;
  Hash: string;
}

export const Block = () => {
  const instance = axios.create({
    baseURL: process.env.REACT_APP_EXPLORER_URL,
  });

  instance.defaults.withCredentials = true;

  const [block, setBlock] = useState<Block>();

  useEffect(() => {
    instance.get(`/block/1`).then((response) => {
      console.log(response);
      const data = response.data;
      setBlock({
        BlockHeight: data.BlockHeight,
        Receipient: data.Receipient,
        Reward: data.Reward,
        Size: data.Size,
        GasUsed: data.GasUsed,
        Hash: data.Hash,
      });
    });
  });

  return (
    <div className="page-block">
      <dl>
        <dt>Height</dt>
        <dd>{block?.BlockHeight}</dd>
        <dt>Receipient</dt>
        <dd>{block?.Receipient}</dd>
        <dt>Reward</dt>
        <dd>{block?.Reward}</dd>
        <dt>Size</dt>
        <dd>{block?.Size}</dd>
        <dt>GasUsed</dt>
        <dd>{block?.GasUsed}</dd>
        <dt>Hash</dt>
        <dd>{block?.Hash}</dd>
      </dl>
    </div>
  );
};
