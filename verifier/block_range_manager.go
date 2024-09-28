package verifier

type eventFetchingBlockRangeManager struct {
	// fields from the configuration
	maxRange uint64

	// internal fields
	lastStart,
	nextStart uint64
}

func newEventFetchingBlockRangeManager(maxRange int, initialStart uint64) *eventFetchingBlockRangeManager {
	return &eventFetchingBlockRangeManager{
		maxRange:  uint64(maxRange),
		lastStart: initialStart,
		nextStart: initialStart,
	}
}

func (m *eventFetchingBlockRangeManager) get(maxEnd uint64) (start, end uint64, skipFetchlog bool) {
	start = m.nextStart

	// Basically, the end block number is the latest block number.
	end = maxEnd

	// Block number is not updated yet.
	if start > end {
		skipFetchlog = true
	}

	// If the range is too wide, divide it into smaller blocks.
	if start < end && end-start > m.maxRange {
		end = start + m.maxRange - 1
	}

	// Update the last start block number.
	m.lastStart = start

	// Next start block is the current end block + 1
	m.nextStart = end + 1

	return
}

func (m *eventFetchingBlockRangeManager) restore() {
	m.nextStart = m.lastStart
}
