import { usePage } from '@inertiajs/react';

export default function Home() {

  const { text } = usePage().props;

  return (
    <div>
      Home: {text}
    </div>
  );
}
