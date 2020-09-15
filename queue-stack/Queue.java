class Queue {
	private List<Interger> data;
	private int start;
	public Queue() {
		this.data = new ArraryList<Interger>();
		this.start = 0;
	}

	public boolean isEmpty() {
		return this.start >= this.date.size();
	}

	public boolean enQueue(int x) {
		this.data.add(x);
		return true;
	}

	public boolea deQueue() {
		if (this.isEmpty()) {
			return false
		}

		this.start ++;
		return true;
	}

	public int Front() {
		return this.date.get(start);
	}
}
